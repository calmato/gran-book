package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	ListFollow(ctx *gin.Context)
	ListFollower(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	Follow(ctx *gin.Context)
	Unfollow(ctx *gin.Context)
}

type userHandler struct {
	userClient user.UserServiceClient
}

func NewUserHandler(uc user.UserServiceClient) UserHandler {
	return &userHandler{
		userClient: uc,
	}
}

// ListFollow - フォロー一覧取得
func (h *userHandler) ListFollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &user.ListFollowRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.userClient.ListFollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollows(out.Follows)
	res := response.NewFollowListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// ListFollower - フォロワー一覧取得
func (h *userHandler) ListFollower(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &user.ListFollowerRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.userClient.ListFollower(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollowers(out.Followers)
	res := response.NewFollowerListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// GetProfile - プロフィール情報取得
func (h *userHandler) GetProfile(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")

	in := &user.GetUserProfileRequest{
		UserId: userID,
	}
	out, err := h.userClient.GetUserProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}

// Follow - フォロー登録
func (h *userHandler) Follow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &user.FollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}
	out, err := h.userClient.Follow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}

// Unfollow - フォロー解除
func (h *userHandler) Unfollow(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &user.UnfollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}
	out, err := h.userClient.Unfollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)
	res := response.NewUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}
