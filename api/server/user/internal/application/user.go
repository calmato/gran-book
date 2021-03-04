package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	GetProfile(ctx context.Context, uid string, cuid string) (*user.User, *output.GetUserProfile, error)
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
	ctx context.Context, uid string, cuid string,
) (*user.User, *output.GetUserProfile, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	followsTotal, followersTotal, err := a.userService.ListFriendsCount(ctx, u)
	if err != nil {
		return nil, nil, err
	}

	isFollow, isFollower, err := a.userService.IsFriend(ctx, u, cuid)
	if err != nil {
		return nil, nil, err
	}

	out := &output.GetUserProfile{
		IsFollow:       isFollow,
		IsFollower:     isFollower,
		FollowsTotal:   followsTotal,
		FollowersTotal: followersTotal,
	}

	return u, out, nil
}
