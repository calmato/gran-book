package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

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
	UID string `json:"sub"`
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

	return fbToken.UID, nil
}

func (r *userRepository) Show(ctx context.Context, uid string) (*user.User, error) {
	u := &user.User{}

	err := r.client.db.First(u, "id = ?", uid).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return u, nil
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

func (r *userRepository) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	uid, err := r.auth.GetUIDByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return uid, nil
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
	payload := strings.Split(token, ".")
	if len(payload) != 3 {
		return nil, xerrors.New("Authorization header is invalid.")
	}

	data, err := base64.RawURLEncoding.DecodeString(payload[1])
	if err != nil {
		return nil, err
	}

	fbToken := &firebaseToken{}

	err = json.Unmarshal(data, fbToken)
	if err != nil {
		return nil, err
	}

	return fbToken, nil
}
