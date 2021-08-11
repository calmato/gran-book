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
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.LIST_LIMIT_DEFAULT))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.LIST_OFFSET_DEFAULT))

	in := &pb.ListFollowRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	out, err := h.userClient.ListFollow(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getFollowListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// ListFollower - フォロワー一覧取得
func (h *userHandler) ListFollower(ctx *gin.Context) {
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.LIST_LIMIT_DEFAULT))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.LIST_OFFSET_DEFAULT))

	in := &pb.ListFollowerRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	out, err := h.userClient.ListFollower(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getFollowerListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// GetProfile - プロフィール情報取得
func (h *userHandler) GetProfile(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &pb.GetUserProfileRequest{
		UserId: userID,
	}

	out, err := h.userClient.GetUserProfile(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserProfileResponse(out)
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

	out, err := h.userClient.Follow(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserProfileResponse(out)
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

	out, err := h.userClient.Unfollow(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserProfileResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) getFollowListResponse(out *pb.FollowListResponse) *response.FollowListResponse {
	users := make([]*response.FollowListUser, len(out.GetFollows()))
	for i, f := range out.GetFollows() {
		user := &response.FollowListUser{
			ID:               f.GetId(),
			Username:         f.GetUsername(),
			ThumbnailURL:     f.GetThumbnailUrl(),
			SelfIntroduction: f.GetSelfIntroduction(),
			IsFollow:         f.GetIsFollow(),
		}

		users[i] = user
	}

	return &response.FollowListResponse{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}

func (h *userHandler) getFollowerListResponse(out *pb.FollowerListResponse) *response.FollowerListResponse {
	users := make([]*response.FollowerListUser, len(out.GetFollowers()))
	for i, f := range out.GetFollowers() {
		user := &response.FollowerListUser{
			ID:               f.GetId(),
			Username:         f.GetUsername(),
			ThumbnailURL:     f.GetThumbnailUrl(),
			SelfIntroduction: f.GetSelfIntroduction(),
			IsFollow:         f.GetIsFollow(),
		}

		users[i] = user
	}

	return &response.FollowerListResponse{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}

func (h *userHandler) getUserProfileResponse(out *pb.UserProfileResponse) *response.UserProfileResponse {
	products := make([]*response.UserProfileProduct, 0)

	return &response.UserProfileResponse{
		ID:               out.GetId(),
		Username:         out.GetUsername(),
		ThumbnailURL:     out.GetThumbnailUrl(),
		SelfIntroduction: out.GetSelfIntroduction(),
		IsFollow:         out.GetIsFollow(),
		IsFollower:       out.GetIsFollower(),
		FollowCount:      out.GetFollowCount(),
		FollowerCount:    out.GetFollowerCount(),
		Rating:           0,
		ReviewCount:      0,
		Products:         products,
	}
}