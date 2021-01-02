package repository

import (
	"context"

	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"github.com/calmato/gran-book/api/user/lib/firebase/authentication"
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

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	_, err := r.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return err
	}

	return r.client.db.Create(&u).Error
}
