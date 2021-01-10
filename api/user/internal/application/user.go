package application

import (
	"context"
	"strings"

	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/calmato/gran-book/api/user/internal/application/validation"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	Authentication(ctx context.Context) (*user.User, error)
	Create(ctx context.Context, in *input.CreateUser) (*user.User, error)
}

type userApplication struct {
	userRequestValidation validation.UserRequestValidation
	userService           user.Service
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(urv validation.UserRequestValidation, us user.Service) UserApplication {
	return &userApplication{
		userRequestValidation: urv,
		userService:           us,
	}
}

func (a *userApplication) Authentication(ctx context.Context) (*user.User, error) {
	u, err := a.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *userApplication) Create(ctx context.Context, in *input.CreateUser) (*user.User, error) {
	err := a.userRequestValidation.CreateUser(in)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username: in.Username,
		Email:    strings.ToLower(in.Email),
		Password: in.Password,
		Gender:   0,
		Role:     0,
	}

	err = a.userService.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
