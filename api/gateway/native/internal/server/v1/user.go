package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler interface {
	ListFollow(ctx *gin.Context)
	ListFollower(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	Follow(ctx *gin.Context)
	Unfollow(ctx *gin.Context)
}

type userHandler struct {
	userClient pb.UserServiceClient
}

func NewUserHandler(userConn *grpc.ClientConn) UserHandler {
	uc := pb.NewUserServiceClient(userConn)

	return &userHandler{
		userClient: uc,
	}
}

// ListFollow - フォロー一覧取得
func (h *userHandler) ListFollow(ctx *gin.Context) {
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &pb.ListFollowRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.ListFollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollows(out.Follows)

	res := h.getFollowListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// ListFollower - フォロワー一覧取得
func (h *userHandler) ListFollower(ctx *gin.Context) {
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	in := &pb.ListFollowerRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.ListFollower(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	fs := entity.NewFollowers(out.Followers)

	res := h.getFollowerListResponse(fs, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// GetProfile - プロフィール情報取得
func (h *userHandler) GetProfile(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &pb.GetUserProfileRequest{
		UserId: userID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.GetUserProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)

	res := h.getUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}

// Follow - フォロー登録
func (h *userHandler) Follow(ctx *gin.Context) {
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &pb.FollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.Follow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)

	res := h.getUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}

// Unfollow - フォロー解除
func (h *userHandler) Unfollow(ctx *gin.Context) {
	userID := ctx.Param("userID")
	followerID := ctx.Param("followerID")

	in := &pb.UnfollowRequest{
		UserId:     userID,
		FollowerId: followerID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.Unfollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	p := entity.NewUserProfile(out.Profile)

	res := h.getUserProfileResponse(p)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) getFollowListResponse(
	fs entity.Follows, limit, offset, total int64,
) *response.FollowListResponse {
	users := make([]*response.FollowListResponse_User, len(fs))
	for i, f := range fs {
		user := &response.FollowListResponse_User{
			ID:               f.Id,
			Username:         f.Username,
			ThumbnailURL:     f.ThumbnailUrl,
			SelfIntroduction: f.SelfIntroduction,
			IsFollow:         f.IsFollow,
		}

		users[i] = user
	}

	return &response.FollowListResponse{
		Users:  users,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

func (h *userHandler) getFollowerListResponse(
	fs entity.Followers, limit, offset, total int64,
) *response.FollowerListResponse {
	users := make([]*response.FollowerListResponse_User, len(fs))
	for i, f := range fs {
		user := &response.FollowerListResponse_User{
			ID:               f.Id,
			Username:         f.Username,
			ThumbnailURL:     f.ThumbnailUrl,
			SelfIntroduction: f.SelfIntroduction,
			IsFollow:         f.IsFollow,
		}

		users[i] = user
	}

	return &response.FollowerListResponse{
		Users:  users,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

func (h *userHandler) getUserProfileResponse(p *entity.UserProfile) *response.UserProfileResponse {
	products := make([]*response.UserProfileResponse_Product, 0)

	return &response.UserProfileResponse{
		ID:               p.Id,
		Username:         p.Username,
		ThumbnailURL:     p.ThumbnailUrl,
		SelfIntroduction: p.SelfIntroduction,
		IsFollow:         p.IsFollow,
		IsFollower:       p.IsFollower,
		FollowCount:      p.FollowCount,
		FollowerCount:    p.FollowerCount,
		Rating:           0,
		ReviewCount:      0,
		Products:         products,
	}
}
