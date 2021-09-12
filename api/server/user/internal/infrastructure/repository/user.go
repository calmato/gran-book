package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/pkg/metadata"
)

type userRepository struct {
	client *database.Client
	auth   *authentication.Auth
}

type firebaseToken struct {
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	UID      string                 `json:"uid,omitempty"`
	Claims   map[string]interface{} `json:"-"`
}

const (
	userTable         = "users"
	relationshipTable = "relationships"
)

var (
	errNotExistsUserInfo            = errors.New("repository: user info is not exists in firebase authentication")
	errNotExistsAuthorizationHeader = errors.New("repository: authorization header is not contain")
	errInvalidAuthorizationHeader   = errors.New("repository: authorization header is invalid")
	errNotVerifyToken               = errors.New("repository: id token is not verify")
)

// NewUserRepository - UserRepositoryの生成
func NewUserRepository(c *database.Client, auth *authentication.Auth) user.Repository {
	return &userRepository{
		client: c,
		auth:   auth,
	}
}

func (r *userRepository) Authentication(ctx context.Context) (string, error) {
	t, err := r.getToken(ctx)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	fbToken, err := r.decodeToken(t)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	return fbToken.Subject, nil
}

func (r *userRepository) List(ctx context.Context, q *database.ListQuery) (user.Users, error) {
	us := user.Users{}

	err := r.client.GetListQuery(userTable, r.client.DB, q).Find(&us).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return us, nil
}

func (r *userRepository) ListFollow(ctx context.Context, q *database.ListQuery) (user.Follows, error) {
	fs := user.Follows{}

	fields := []string{
		"relationships.follow_id",
		"relationships.follower_id",
		"users.username",
		"users.thumbnail_url",
		"users.self_introduction",
	}
	err := r.client.
		GetListQuery(relationshipTable, r.client.DB, q).
		Select(fields).
		Joins("LEFT JOIN users ON relationships.follow_id = users.id").
		Find(&fs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return fs, nil
}

func (r *userRepository) ListFollower(ctx context.Context, q *database.ListQuery) (user.Followers, error) {
	fs := user.Followers{}

	fields := []string{
		"relationships.follow_id",
		"relationships.follower_id",
		"users.username",
		"users.thumbnail_url",
		"users.self_introduction",
	}

	err := r.client.
		GetListQuery(relationshipTable, r.client.DB, q).
		Select(fields).
		Joins("LEFT JOIN users ON relationships.follower_id = users.id").
		Find(&fs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return fs, nil
}

func (r *userRepository) ListInstanceID(ctx context.Context, q *database.ListQuery) ([]string, error) {
	instanceIDs := []string{}

	err := r.client.
		GetListQuery(userTable, r.client.DB, q).
		Select("instance_id").
		Find(&instanceIDs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return instanceIDs, nil
}

func (r *userRepository) ListFollowID(ctx context.Context, followerID string, followIDs ...string) ([]string, error) {
	ids := []string{}

	sql := r.client.DB.
		Table(relationshipTable).
		Select("follow_id").
		Where("follower_id = ?", followerID)

	if len(followIDs) > 0 {
		sql = sql.Where("follow_id IN (?)", followIDs)
	}

	err := sql.Find(&ids).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return ids, nil
}

func (r *userRepository) ListFollowerID(ctx context.Context, followID string, followerIDs ...string) ([]string, error) {
	ids := []string{}

	sql := r.client.DB.
		Table(relationshipTable).
		Select("follower_id").
		Where("follow_id = ?", followID)

	if len(followerIDs) > 0 {
		sql = sql.Where("follower_id IN (?)", followerIDs)
	}

	err := sql.Find(&ids).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return ids, nil
}

func (r *userRepository) Count(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(userTable, r.client.DB, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *userRepository) CountRelationship(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(relationshipTable, r.client.DB, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *userRepository) MultiGet(ctx context.Context, userIDs []string) (user.Users, error) {
	us := user.Users{}

	err := r.client.DB.Table(userTable).Where("id IN (?)", userIDs).Find(&us).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return us, nil
}

func (r *userRepository) Get(ctx context.Context, userID string) (*user.User, error) {
	u := &user.User{}

	err := r.client.DB.Table(userTable).First(u, "id = ?", userID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return u, nil
}

func (r *userRepository) GetAdmin(ctx context.Context, userID string) (*user.User, error) {
	u := &user.User{}

	err := r.client.DB.
		Table(userTable).
		Where("role IN (?)", []int{user.AdminRole, user.DeveloperRole, user.OperatorRole}).
		First(u, "id = ?", userID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return u, nil
}

func (r *userRepository) GetRelationship(
	ctx context.Context, followID string, followerID string,
) (*user.Relationship, error) {
	rs := &user.Relationship{}

	err := r.client.DB.
		Table(relationshipTable).
		First(rs, "follow_id = ? AND follower_id = ?", followID, followerID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rs, nil
}

func (r *userRepository) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := r.auth.GetUIDByEmail(ctx, email)
	if err != nil {
		return "", exception.NotFound.New(err)
	}

	return uid, nil
}

func (r *userRepository) GetRelationshipIDByUserID(
	ctx context.Context, followID string, followerID string,
) (int, error) {
	rs := &user.Relationship{}

	err := r.client.DB.
		Table(relationshipTable).
		Select("id").
		First(rs, "follow_id = ? AND follower_id = ?", followID, followerID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return rs.ID, nil
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	err = r.client.DB.Table(userTable).Create(&u).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) CreateWithOAuth(ctx context.Context, u *user.User) error {
	au, err := r.getAuth(ctx, u.ID)
	if err != nil {
		return exception.Unauthorized.New(err)
	}

	u.Username = r.getUsername(au.UserInfo.DisplayName, u.ID)
	u.Email = au.UserInfo.Email
	u.ThumbnailURL = au.UserInfo.PhotoURL

	err = r.client.DB.Table(userTable).Create(&u).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) CreateRelationship(ctx context.Context, rs *user.Relationship) error {
	err := r.client.DB.Table(relationshipTable).Create(&rs).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	au, err := r.getAuth(ctx, u.ID)
	if err != nil {
		return exception.Unauthorized.New(err)
	}

	if u.Email != au.UserInfo.Email {
		err = r.auth.UpdateEmail(ctx, u.ID, u.Email)
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	err = r.client.DB.Table(userTable).Save(u).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, uid string, password string) error {
	err := r.auth.UpdatePassword(ctx, uid, password)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, userID string) error {
	err := r.client.DB.Table(userTable).Where("id = ?", userID).Delete(&user.User{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	err = r.auth.DeleteUser(ctx, userID)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) DeleteRelationship(ctx context.Context, relationshipID int) error {
	err := r.client.DB.Table(relationshipTable).Where("id = ?", relationshipID).Delete(&user.Relationship{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) getAuth(ctx context.Context, userID string) (*auth.UserRecord, error) {
	au, err := r.auth.GetUserByUID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if au.UserInfo == nil {
		return nil, errNotExistsUserInfo
	}

	return au, nil
}

func (r *userRepository) getToken(ctx context.Context) (string, error) {
	authorization, err := metadata.Get(ctx, "Authorization")
	if err != nil {
		return "", err
	}

	if authorization == "" {
		return "", errNotExistsAuthorizationHeader
	}

	t := strings.Replace(authorization, "Bearer ", "", 1)
	return t, nil
}

func (r *userRepository) decodeToken(token string) (*firebaseToken, error) {
	s := strings.Split(token, ".")
	if len(s) != 3 {
		return nil, errInvalidAuthorizationHeader
	}

	data, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return nil, err
	}

	fbToken := &firebaseToken{}

	err = json.Unmarshal(data, fbToken)
	if err != nil {
		return nil, err
	}

	err = r.verifyToken(fbToken)
	if err != nil {
		return nil, err
	}

	return fbToken, nil
}

func (r *userRepository) verifyToken(t *firebaseToken) error {
	now := time.Now().Unix()

	if t.IssuedAt > now {
		return fmt.Errorf("%w, issued at future timestamp: %d", errNotVerifyToken, t.IssuedAt)
	}

	if t.Expires < now {
		return fmt.Errorf("%w, has expired. expired at: %d", errNotVerifyToken, t.Expires)
	}

	if t.Subject == "" {
		return fmt.Errorf("%w, has empty subject claim", errNotVerifyToken)
	}

	if len(t.Subject) > 128 {
		return fmt.Errorf("%w, has a subject claim longer than 128 characters", errNotVerifyToken)
	}

	return nil
}

// OAuth認証による初回User登録時、UIDの先頭12文字を取得して作成
// e.g.) 12345678-qwer-asdf-zxcv-uiophjklvbnm -> 12345678qwer
func (r *userRepository) getUsername(displayName string, uid string) string {
	if displayName != "" {
		return displayName
	}

	str := strings.Replace(uid, "-", "", -1)
	return fmt.Sprintf("user-%s", str[0:12])
}
