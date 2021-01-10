package application

import (
	"context"
	"strings"

	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/calmato/gran-book/api/user/internal/application/validation"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
)

// AuthApplication - Authアプリケーションのインターフェース
type AuthApplication interface {
	Authentication(ctx context.Context) (*user.User, error)
	Create(ctx context.Context, in *input.CreateAuth) (*user.User, error)
	UpdatePassword(ctx context.Context, in *input.UpdateAuthPassword, u *user.User) error
}

type authApplication struct {
	authRequestValidation validation.AuthRequestValidation
	userService           user.Service
}

// NewAuthApplication - AuthApplicationの生成
func NewAuthApplication(urv validation.AuthRequestValidation, us user.Service) AuthApplication {
	return &authApplication{
		authRequestValidation: urv,
		userService:           us,
	}
}

func (a *authApplication) Authentication(ctx context.Context) (*user.User, error) {
	u, err := a.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *authApplication) Create(ctx context.Context, in *input.CreateAuth) (*user.User, error) {
	err := a.authRequestValidation.CreateAuth(in)
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

func (a *authApplication) UpdatePassword(ctx context.Context, in *input.UpdateAuthPassword, u *user.User) error {
	err := a.authRequestValidation.UpdateAuthPassword(in)
	if err != nil {
		return err
	}

	return a.userService.UpdatePassword(ctx, u.ID, in.Password)
}
