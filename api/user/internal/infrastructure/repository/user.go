package repository

import (
	"context"
	"strings"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"github.com/calmato/gran-book/api/user/lib/firebase/authentication"
	"github.com/calmato/gran-book/api/user/lib/metadata"
	"golang.org/x/xerrors"
)

type userRepository struct {
	client *Client
	auth   *authentication.Auth
}

// NewUserRepository - UserRepositoryの生成
func NewUserRepository(c *Client, auth *authentication.Auth) user.Repository {
	return &userRepository{
		client: c,
		auth:   auth,
	}
}

func (r *userRepository) Authentication(ctx context.Context) (*user.User, error) {
	t, err := getToken(ctx)
	if err != nil {
		return nil, exception.Unauthorized.New(err)
	}

	uid, err := r.auth.VerifyIDToken(ctx, t)
	if err != nil {
		return nil, exception.Unauthorized.New(err)
	}

	u := &user.User{}

	err = r.client.db.First(u, "id = ?", uid).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
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
