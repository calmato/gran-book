package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/admin/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
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

func (h *userHandler) List(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	field := ctx.DefaultQuery("field", "")
	value := ctx.DefaultQuery("value", "")
	by := ctx.DefaultQuery("by", "")
	direction := ctx.DefaultQuery("direction", "")

	in := &pb.ListUserRequest{
		Limit:  limit,
		Offset: offset,
	}

	if field != "" {
		search := &pb.Search{
			Field: field,
			Value: value,
		}

		in.Search = search
	}

	// TODO: CamelCase -> snake_case に変換する関数作成したい..
	if by != "" {
		order := &pb.Order{
			Field:   by,
			OrderBy: entity.OrderBy(0).Value(direction).Proto(),
		}

		in.Order = order
	}

	out, err := h.userClient.ListUser(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &pb.GetUserRequest{
		UserId: userID,
	}

	out, err := h.userClient.GetUser(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) getUserResponse(out *pb.UserResponse) *response.UserResponse {
	return &response.UserResponse{
		ID:               out.GetId(),
		Username:         out.GetUsername(),
		Email:            out.GetEmail(),
		PhoneNumber:      out.GetPhoneNumber(),
		ThumbnailURL:     out.GetThumbnailUrl(),
		SelfIntroduction: out.GetSelfIntroduction(),
		LastName:         out.GetLastName(),
		FirstName:        out.GetFirstName(),
		LastNameKana:     out.GetLastNameKana(),
		FirstNameKana:    out.GetFirstNameKana(),
		CreatedAt:        out.GetCreatedAt(),
		UpdatedAt:        out.GetUpdatedAt(),
	}
}

func (h *userHandler) getUserListResponse(out *pb.UserListResponse) *response.UserListResponse {
	users := make([]*response.UserListUser, len(out.GetUsers()))
	for i, u := range out.GetUsers() {
		user := &response.UserListUser{
			ID:               u.GetId(),
			Username:         u.GetUsername(),
			Email:            u.GetEmail(),
			PhoneNumber:      u.GetPhoneNumber(),
			ThumbnailURL:     u.GetThumbnailUrl(),
			SelfIntroduction: u.GetSelfIntroduction(),
			LastName:         u.GetLastName(),
			FirstName:        u.GetFirstName(),
			LastNameKana:     u.GetLastNameKana(),
			FirstNameKana:    u.GetFirstNameKana(),
			CreatedAt:        u.GetCreatedAt(),
			UpdatedAt:        u.GetUpdatedAt(),
		}

		users[i] = user
	}

	return &response.UserListResponse{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}
