package application

import (
	"context"

	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	Create(ctx context.Context, in *input.CreateUser) (*user.User, error)
}

type userApplication struct {
	userService user.Service
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(us user.Service) UserApplication {
	return &userApplication{
		userService: us,
	}
}

func (a *userApplication) Create(ctx context.Context, in *input.CreateUser) (*user.User, error) {
	// TODO: validation

	u := &user.User{
		Username: in.Username,
		Email:    in.Email,
		Password: in.Password,
		Gender:   0,
		Role:     0,
	}

	err := a.userService.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
