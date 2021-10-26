package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/entity"
	response "github.com/calmato/gran-book/api/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/gin-gonic/gin"
)

// listUserFollow - フォロー一覧取得
func (h *apiV1Handler) listUserFollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	fs, total, err := h.userListFollow(c, userID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.FollowListResponse{
		Users:  entity.NewFollows(fs),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	ctx.JSON(http.StatusOK, res)
}

// listUserFollower - フォロワー一覧取得
func (h *apiV1Handler) listUserFollower(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	fs, total, err := h.userListFollower(c, userID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.FollowerListResponse{
		Users:  entity.NewFollowers(fs),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	ctx.JSON(http.StatusOK, res)
}

// getUserProfile - プロフィール情報取得
func (h *apiV1Handler) getUserProfile(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	p, err := h.userGetUserProfile(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.UserProfileResponse{
		UserProfile: entity.NewUserProfile(p),
	}
	ctx.JSON(http.StatusOK, res)
}

// userFollow - フォロー登録
func (h *apiV1Handler) userFollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &user.FollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}
	out, err := h.User.Follow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	p := gentity.NewUserProfile(out.Profile)

	res := &response.UserProfileResponse{
		UserProfile: entity.NewUserProfile(p),
	}
	ctx.JSON(http.StatusOK, res)
}

// userUnfollow - フォロー解除
func (h *apiV1Handler) userUnfollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &user.UnfollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}
	out, err := h.User.Unfollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	p := gentity.NewUserProfile(out.Profile)

	res := &response.UserProfileResponse{
		UserProfile: entity.NewUserProfile(p),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) userMultiGetUser(ctx context.Context, userIDs []string) (gentity.Users, error) {
	in := &user.MultiGetUserRequest{
		UserIds: userIDs,
	}
	out, err := h.User.MultiGetUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewUsers(out.Users), nil
}

func (h *apiV1Handler) userGetUser(ctx context.Context, userID string) (*gentity.User, error) {
	in := &user.GetUserRequest{
		UserId: userID,
	}
	out, err := h.User.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.User == nil {
		err := fmt.Errorf("user is not found: %s", userID)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewUser(out.User), nil
}

func (h *apiV1Handler) userListFollow(
	ctx context.Context, userID string, limit int64, offset int64,
) (gentity.Follows, int64, error) {
	in := &user.ListFollowRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.User.ListFollow(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewFollows(out.Follows), out.Total, nil
}

func (h *apiV1Handler) userListFollower(
	ctx context.Context, userID string, limit int64, offset int64,
) (gentity.Followers, int64, error) {
	in := &user.ListFollowerRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.User.ListFollower(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewFollowers(out.Followers), out.Total, nil
}

func (h *apiV1Handler) userGetUserProfile(ctx context.Context, userID string) (*gentity.UserProfile, error) {
	in := &user.GetUserProfileRequest{
		UserId: userID,
	}
	out, err := h.User.GetUserProfile(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewUserProfile(out.Profile), nil
}
