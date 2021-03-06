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
	GetUserProfile(ctx context.Context, uid string, cuid string) (*user.User, *output.UserProfile, error)
	RegisterFollow(ctx context.Context, uid string, cuid string) (*user.User, *output.UserProfile, error)
	UnregisterFollow(ctx context.Context, id int64) (*user.User, *output.UserProfile, error)
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

func (a *userApplication) GetUserProfile(
	ctx context.Context, uid string, cuid string,
) (*user.User, *output.UserProfile, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, uid)
	isFollow, isFollower := a.userService.IsFriend(ctx, uid, cuid)

	out := &output.UserProfile{
		IsFollow:      isFollow,
		IsFollower:    isFollower,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}

	return u, out, nil
}

func (a *userApplication) RegisterFollow(
	ctx context.Context, uid string, cuid string,
) (*user.User, *output.UserProfile, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	r := &user.Relationship{
		FollowID:   cuid,
		FollowerID: uid,
	}

	err = a.userService.CreateRelationship(ctx, r)
	if err != nil {
		return nil, nil, err
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, uid)
	isFollow, isFollower := a.userService.IsFriend(ctx, uid, cuid)

	out := &output.UserProfile{
		IsFollow:      isFollow,
		IsFollower:    isFollower,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}

	return u, out, nil
}

func (a *userApplication) UnregisterFollow(ctx context.Context, id int64) (*user.User, *output.UserProfile, error) {
	r, err := a.userService.ShowRelationship(ctx, id)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	u, err := a.userService.Show(ctx, r.FollowerID)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	err = a.userService.DeleteRelationship(ctx, r.ID)
	if err != nil {
		return nil, nil, err
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, r.FollowerID)
	isFollow, isFollower := a.userService.IsFriend(ctx, r.FollowerID, r.FollowID)

	out := &output.UserProfile{
		IsFollow:      isFollow,
		IsFollower:    isFollower,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}

	return u, out, nil
}
