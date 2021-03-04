package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	GetProfile(ctx context.Context, in *input.GetUserProfile) (*user.User, *output.GetUserProfile, error)
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

func (a *userApplication) GetProfile(
	ctx context.Context, in *input.GetUserProfile,
) (*user.User, *output.GetUserProfile, error) {
	err := a.userRequestValidation.GetUserProfile(in)
	if err != nil {
		return nil, nil, err
	}

	u, err := a.userService.Show(ctx, in.ID)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	followsTotal, followersTotal, err := a.userService.ListFriendsCount(ctx, u)
	if err != nil {
		return nil, nil, err
	}

	out := &output.GetUserProfile{
		FollowsTotal:   followsTotal,
		FollowersTotal: followersTotal,
	}

	return u, out, nil
}
