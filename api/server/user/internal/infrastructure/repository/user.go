package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/lib/metadata"
	"golang.org/x/xerrors"
)

type userRepository struct {
	client *Client
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
func NewUserRepository(c *Client, auth *authentication.Auth) user.Repository {
	return &userRepository{
		client: c,
		auth:   auth,
	}
}

func (r *userRepository) Authentication(ctx context.Context) (string, error) {
	t, err := getToken(ctx)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	fbToken, err := decodeToken(t)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	return fbToken.Subject, nil
}

func (r *userRepository) List(ctx context.Context, q *domain.ListQuery) ([]*user.User, error) {
	us := []*user.User{}

	sql := r.client.db
	db := r.client.getListQuery(sql, q)

	err := db.Find(&us).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return us, nil
}

func (r *userRepository) ListFollow(ctx context.Context, q *domain.ListQuery) ([]*user.Follow, error) {
	fs := []*user.Follow{}

	columns := []string{
		"relationships.follow_id",
		"relationships.follower_id",
		"users.username",
		"users.thumbnail_url",
		"users.self_introduction",
	}
	sql := r.client.db.
		Table("relationships").
		Select(strings.Join(columns, ", ")).
		Joins("LEFT JOIN users ON relationships.follower_id = users.id")
	db := r.client.getListQuery(sql, q)

	err := db.Scan(&fs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return fs, nil
}

func (r *userRepository) ListFollower(ctx context.Context, q *domain.ListQuery) ([]*user.Follower, error) {
	fs := []*user.Follower{}

	columns := []string{
		"relationships.follow_id",
		"relationships.follower_id",
		"users.username",
		"users.thumbnail_url",
		"users.self_introduction",
	}
	sql := r.client.db.
		Table("relationships").
		Select(strings.Join(columns, ", ")).
		Joins("LEFT JOIN users ON relationships.follow_id = users.id")
	db := r.client.getListQuery(sql, q)

	err := db.Scan(&fs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return fs, nil
}

func (r *userRepository) ListCount(ctx context.Context, q *domain.ListQuery) (int64, error) {
	sql := r.client.db.Table("users")
	return r.client.getListCount(sql, q)
}

func (r *userRepository) ListRelationshipCount(ctx context.Context, q *domain.ListQuery) (int64, error) {
	sql := r.client.db.Table("relationships")
	return r.client.getListCount(sql, q)
}

func (r *userRepository) Show(ctx context.Context, uid string) (*user.User, error) {
	u := &user.User{}

	err := r.client.db.First(u, "id = ?", uid).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return u, nil
}

func (r *userRepository) ShowRelationship(ctx context.Context, id int64) (*user.Relationship, error) {
	rs := &user.Relationship{}

	err := r.client.db.First(rs, "id = ?", id).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rs, nil
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	err = r.client.db.Create(&u).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) CreateRelationship(ctx context.Context, rs *user.Relationship) error {
	err := r.client.db.Create(&rs).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	au, err := r.auth.GetUserByUID(ctx, u.ID)
	if err != nil {
		return exception.NotFound.New(err)
	}

	if au.UserInfo == nil {
		err = xerrors.New("UserInfo is not exists in Firebase Authentication")
		return exception.ErrorInDatastore.New(err)
	}

	if u.Email != au.UserInfo.Email {
		err = r.auth.UpdateEmail(ctx, u.ID, u.Email)
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	err = r.client.db.Save(u).Error
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

func (r *userRepository) DeleteRelationship(ctx context.Context, id int64) error {
	err := r.client.db.Delete(&user.Relationship{}, id).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *userRepository) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := r.auth.GetUIDByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (r *userRepository) GetRelationshipIDByUID(
	ctx context.Context, followID string, followerID string,
) (int64, error) {
	rs := &user.Relationship{}

	err := r.client.db.Select("id").First(rs, "follow_id = ? AND follower_id = ?", followID, followerID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return rs.ID, nil
}

func getToken(ctx context.Context) (string, error) {
	authorization, err := metadata.Get(ctx, "Authorization")
	if err != nil {
		return "", err
	}

	if authorization == "" {
		return "", xerrors.New("Authorization header is not contain.")
	}

	t := strings.Replace(authorization, "Bearer ", "", 1)
	return t, nil
}

func decodeToken(token string) (*firebaseToken, error) {
	s := strings.Split(token, ".")
	if len(s) != 3 {
		return nil, xerrors.New("Authorization header is invalid.")
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

	err = verifyToken(fbToken)
	if err != nil {
		return nil, err
	}

	return fbToken, nil
}

func verifyToken(t *firebaseToken) error {
	now := time.Now().Unix()

	verifyTokenMsg := "See https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to " +
		"retrieve a valid ID token."

	if t.IssuedAt > now {
		return xerrors.Errorf("ID token issued at future timestamp: %d", t.IssuedAt)
	}

	if t.Expires < now {
		return xerrors.Errorf("ID token has expired. Expired at: %d", t.Expires)
	}

	if t.Subject == "" {
		return xerrors.Errorf("ID token has empty 'sub' (subject) claim. %s", verifyTokenMsg)
	}

	if len(t.Subject) > 128 {
		return xerrors.Errorf("ID token has a 'sub' (subject) claim longer than 128 characters. %s", verifyTokenMsg)
	}

	return nil
}
