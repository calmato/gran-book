package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/pkg/metadata"
	"gorm.io/gorm"
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
		return "", toDBError(err)
	}

	fbToken, err := r.decodeToken(t)
	if err != nil {
		return "", toDBError(err)
	}

	return fbToken.Subject, nil
}

func (r *userRepository) List(ctx context.Context, q *database.ListQuery) (user.Users, error) {
	us := user.Users{}

	err := r.client.GetListQuery(userTable, r.client.DB, q).Find(&us).Error
	return us, toDBError(err)
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
	return fs, toDBError(err)
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
	return fs, toDBError(err)
}

func (r *userRepository) ListInstanceID(ctx context.Context, q *database.ListQuery) ([]string, error) {
	instanceIDs := []string{}

	err := r.client.
		GetListQuery(userTable, r.client.DB, q).
		Select("instance_id").
		Find(&instanceIDs).Error
	return instanceIDs, toDBError(err)
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
	return ids, toDBError(err)
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
	return ids, toDBError(err)
}

func (r *userRepository) Count(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(userTable, r.client.DB, q)
	return total, toDBError(err)
}

func (r *userRepository) CountRelationship(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(relationshipTable, r.client.DB, q)
	return total, toDBError(err)
}

func (r *userRepository) MultiGet(ctx context.Context, userIDs []string) (user.Users, error) {
	us := user.Users{}

	err := r.client.DB.Table(userTable).Where("id IN (?)", userIDs).Find(&us).Error
	return us, toDBError(err)
}

func (r *userRepository) Get(ctx context.Context, userID string) (*user.User, error) {
	u := &user.User{}

	err := r.client.DB.Table(userTable).First(u, "id = ?", userID).Error
	return u, toDBError(err)
}

func (r *userRepository) GetAdmin(ctx context.Context, userID string) (*user.User, error) {
	u := &user.User{}

	err := r.client.DB.
		Table(userTable).
		Where("role IN (?)", []int{user.AdminRole, user.DeveloperRole, user.OperatorRole}).
		First(u, "id = ?", userID).Error
	return u, toDBError(err)
}

func (r *userRepository) GetRelationship(
	ctx context.Context, followID string, followerID string,
) (*user.Relationship, error) {
	rs := &user.Relationship{}

	err := r.client.DB.
		Table(relationshipTable).
		First(rs, "follow_id = ? AND follower_id = ?", followID, followerID).Error
	return rs, toDBError(err)
}

func (r *userRepository) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := r.auth.GetUIDByEmail(ctx, email)
	return uid, toFirebaseError(err)
}

func (r *userRepository) GetRelationshipIDByUserID(
	ctx context.Context, followID string, followerID string,
) (int, error) {
	rs := &user.Relationship{}

	err := r.client.DB.
		Table(relationshipTable).
		Select("id").
		First(rs, "follow_id = ? AND follower_id = ?", followID, followerID).Error
	return rs.ID, toDBError(err)
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return toFirebaseError(err)
	}

	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createUser(tx, u)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	return toDBError(tx.Commit().Error)
}

func (r *userRepository) CreateWithOAuth(ctx context.Context, u *user.User) error {
	au, err := r.getAuth(ctx, u.ID)
	if err != nil {
		return toDBError(err)
	}

	u.Username = r.getUsername(au.UserInfo.DisplayName, u.ID)
	u.Email = au.UserInfo.Email
	u.ThumbnailURL = au.UserInfo.PhotoURL

	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createUser(tx, u)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	return toDBError(tx.Commit().Error)
}

func (r *userRepository) createUser(tx *gorm.DB, u *user.User) error {
	return tx.Table(userTable).Create(&u).Error
}

func (r *userRepository) CreateRelationship(ctx context.Context, rs *user.Relationship) error {
	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createRelationship(tx, rs)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	return toDBError(tx.Commit().Error)
}

func (r *userRepository) createRelationship(tx *gorm.DB, rs *user.Relationship) error {
	return tx.Table(relationshipTable).Create(&rs).Error
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	au, err := r.getAuth(ctx, u.ID)
	if err != nil {
		return toDBError(err)
	}

	if u.Email != au.UserInfo.Email {
		err = r.auth.UpdateEmail(ctx, u.ID, u.Email)
		if err != nil {
			return toFirebaseError(err)
		}
	}

	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.updateUser(tx, u)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	return toDBError(tx.Commit().Error)
}

func (r *userRepository) updateUser(tx *gorm.DB, u *user.User) error {
	return tx.Table(userTable).Save(&u).Error
}

func (r *userRepository) UpdatePassword(ctx context.Context, uid string, password string) error {
	err := r.auth.UpdatePassword(ctx, uid, password)
	return toFirebaseError(err)
}

func (r *userRepository) Delete(ctx context.Context, userID string) error {
	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.deleteUser(tx, userID)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	err = r.auth.DeleteUser(ctx, userID)
	return toFirebaseError(err)
}

func (r *userRepository) deleteUser(tx *gorm.DB, userID string) error {
	return tx.Table(userTable).Where("id = ?", userID).Delete(&user.User{}).Error
}

func (r *userRepository) DeleteRelationship(ctx context.Context, relationshipID int) error {
	tx, err := r.client.Begin()
	if err != nil {
		return toDBError(err)
	}
	defer r.client.Close(tx)

	err = r.deleteRelationship(tx, relationshipID)
	if err != nil {
		tx.Rollback()
		return toDBError(err)
	}

	return toDBError(tx.Commit().Error)
}

func (r *userRepository) deleteRelationship(tx *gorm.DB, relationshipID int) error {
	return tx.Table(relationshipTable).Where("id = ?", relationshipID).Delete(&user.Relationship{}).Error
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
