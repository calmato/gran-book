package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

// listUserFollow - フォロー一覧取得
func (h *apiV1Handler) listUserFollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &user.ListFollowRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.User.ListFollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollows(out.Follows)
	res := response.NewFollowListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// listUserFollower - フォロワー一覧取得
func (h *apiV1Handler) listUserFollower(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &user.ListFollowerRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.User.ListFollower(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollowers(out.Followers)
	res := response.NewFollowerListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// getUserProfile - プロフィール情報取得
func (h *apiV1Handler) getUserProfile(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")

	in := &user.GetUserProfileRequest{
		UserId: userID,
	}
	out, err := h.User.GetUserProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
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

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
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

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}
