package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
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

	res := h.getFollowListResponse(out)
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

	res := h.getFollowerListResponse(out)
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

	c := util.SetMetadata(ctx)
	out, err := h.userClient.Follow(c, in)
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

	c := util.SetMetadata(ctx)
	out, err := h.userClient.Unfollow(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserProfileResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) getFollowListResponse(out *pb.FollowListResponse) *pb.FollowListV1Response {
	users := make([]*pb.FollowListV1Response_User, len(out.GetFollows()))
	for i, f := range out.GetFollows() {
		user := &pb.FollowListV1Response_User{
			Id:               f.GetId(),
			Username:         f.GetUsername(),
			ThumbnailUrl:     f.GetThumbnailUrl(),
			SelfIntroduction: f.GetSelfIntroduction(),
			IsFollow:         f.GetIsFollow(),
		}

		users[i] = user
	}

	return &pb.FollowListV1Response{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}

func (h *userHandler) getFollowerListResponse(out *pb.FollowerListResponse) *pb.FollowerListV1Response {
	users := make([]*pb.FollowerListV1Response_User, len(out.GetFollowers()))
	for i, f := range out.GetFollowers() {
		user := &pb.FollowerListV1Response_User{
			Id:               f.GetId(),
			Username:         f.GetUsername(),
			ThumbnailUrl:     f.GetThumbnailUrl(),
			SelfIntroduction: f.GetSelfIntroduction(),
			IsFollow:         f.GetIsFollow(),
		}

		users[i] = user
	}

	return &pb.FollowerListV1Response{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}

func (h *userHandler) getUserProfileResponse(out *pb.UserProfileResponse) *pb.UserProfileV1Response {
	products := make([]*pb.UserProfileV1Response_Product, 0)

	return &pb.UserProfileV1Response{
		Id:               out.GetId(),
		Username:         out.GetUsername(),
		ThumbnailUrl:     out.GetThumbnailUrl(),
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
