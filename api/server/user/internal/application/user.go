package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	ListFollow(
		ctx context.Context, in *input.ListFollow, uid string, cuid string,
	) ([]*user.Follow, *output.ListQuery, error)
	ListFollower(
		ctx context.Context, in *input.ListFollower, uid string, cuid string,
	) ([]*user.Follower, *output.ListQuery, error)
	Show(ctx context.Context, uid string) (*user.User, error)
	GetUserProfile(ctx context.Context, uid string, cuid string) (*user.User, *output.UserProfile, error)
	RegisterFollow(ctx context.Context, uid string, cuid string) (*user.User, *output.UserProfile, error)
	UnregisterFollow(ctx context.Context, uid string, cuid string) (*user.User, *output.UserProfile, error)
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

func (a *userApplication) ListFollow(
	ctx context.Context, in *input.ListFollow, uid string, cuid string,
) ([]*user.Follow, *output.ListQuery, error) {
	err := a.userRequestValidation.ListFollow(in)
	if err != nil {
		return nil, nil, err
	}

	q := &domain.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    uid,
			},
		},
	}

	if in.By != "" {
		o := &domain.QueryOrder{
			By:        in.By,
			Direction: in.Direction,
		}

		q.Order = o
	}

	fs, err := a.userService.ListFollow(ctx, q, cuid)
	if err != nil {
		return nil, nil, err
	}

	total, _, err := a.userService.ListFriendCount(ctx, uid)
	if err != nil {
		return nil, nil, err
	}

	out := &output.ListQuery{
		Limit:  q.Limit,
		Offset: q.Offset,
		Total:  total,
	}

	if q.Order != nil {
		o := &output.QueryOrder{
			By:        q.Order.By,
			Direction: q.Order.Direction,
		}

		out.Order = o
	}

	return fs, out, nil
}

func (a *userApplication) ListFollower(
	ctx context.Context, in *input.ListFollower, uid string, cuid string,
) ([]*user.Follower, *output.ListQuery, error) {
	err := a.userRequestValidation.ListFollower(in)
	if err != nil {
		return nil, nil, err
	}

	q := &domain.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    uid,
			},
		},
	}

	if in.By != "" {
		o := &domain.QueryOrder{
			By:        in.By,
			Direction: in.Direction,
		}

		q.Order = o
	}

	fs, err := a.userService.ListFollower(ctx, q, cuid)
	if err != nil {
		return nil, nil, err
	}

	_, total, err := a.userService.ListFriendCount(ctx, uid)
	if err != nil {
		return nil, nil, err
	}

	out := &output.ListQuery{
		Limit:  q.Limit,
		Offset: q.Offset,
		Total:  total,
	}

	if q.Order != nil {
		o := &output.QueryOrder{
			By:        q.Order.By,
			Direction: q.Order.Direction,
		}

		out.Order = o
	}

	return fs, out, nil
}

func (a *userApplication) Show(ctx context.Context, uid string) (*user.User, error) {
	return a.userService.Show(ctx, uid)
}

func (a *userApplication) GetUserProfile(
	ctx context.Context, uid string, cuid string,
) (*user.User, *output.UserProfile, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, uid)
	if err != nil {
		return nil, nil, err
	}

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

	err = a.userService.ValidationRelationship(ctx, r)
	if err != nil {
		return nil, nil, err
	}

	err = a.userService.CreateRelationship(ctx, r)
	if err != nil {
		return nil, nil, err
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, uid)
	if err != nil {
		return nil, nil, err
	}

	isFollow, isFollower := a.userService.IsFriend(ctx, uid, cuid)

	out := &output.UserProfile{
		IsFollow:      isFollow,
		IsFollower:    isFollower,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}

	return u, out, nil
}

func (a *userApplication) UnregisterFollow(
	ctx context.Context, uid string, cuid string,
) (*user.User, *output.UserProfile, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	r, err := a.userService.ShowRelationshipByUID(ctx, uid, cuid)
	if err != nil {
		return nil, nil, exception.NotFound.New(err)
	}

	err = a.userService.DeleteRelationship(ctx, r.ID)
	if err != nil {
		return nil, nil, err
	}

	followCount, followerCount, err := a.userService.ListFriendCount(ctx, r.FollowerID)
	if err != nil {
		return nil, nil, err
	}

	isFollow, isFollower := a.userService.IsFriend(ctx, r.FollowerID, r.FollowID)

	out := &output.UserProfile{
		IsFollow:      isFollow,
		IsFollower:    isFollower,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}

	return u, out, nil
}
