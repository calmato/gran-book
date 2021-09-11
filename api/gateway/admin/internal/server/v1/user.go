package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/admin/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/conv"
	"github.com/calmato/gran-book/api/gateway/admin/proto/user"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type userHandler struct {
	userClient user.UserServiceClient
}

func NewUserHandler(uc user.UserServiceClient) UserHandler {
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

	in := &user.ListUserRequest{
		Limit:  limit,
		Offset: offset,
	}

	if field != "" {
		search := &user.Search{
			Field: field,
			Value: value,
		}

		in.Search = search
	}

	if by != "" {
		// TODO: キャメルケース->スネークケースの変換はマイクロサービスでやる
		field, err := conv.CamelToSnake(field)
		if err != nil {
			util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
			return
		}

		order := &user.Order{
			Field:   field,
			OrderBy: entity.OrderBy(0).Value(direction).Proto(),
		}

		in.Order = order
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.ListUser(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(out.Users)

	res := h.getUserListResponse(us, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &user.GetUserRequest{
		UserId: userID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.userClient.GetUser(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewUser(out.User)

	res := h.getUserResponse(u)
	ctx.JSON(http.StatusOK, res)
}

func (h *userHandler) getUserResponse(u *entity.User) *response.UserResponse {
	return &response.UserResponse{
		ID:               u.Id,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		ThumbnailURL:     u.ThumbnailUrl,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

func (h *userHandler) getUserListResponse(us entity.Users, limit, offset, total int64) *response.UserListResponse {
	users := make([]*response.UserListUser, len(us))
	for i, u := range us {
		user := &response.UserListUser{
			ID:               u.Id,
			Username:         u.Username,
			Email:            u.Email,
			PhoneNumber:      u.PhoneNumber,
			ThumbnailURL:     u.ThumbnailUrl,
			SelfIntroduction: u.SelfIntroduction,
			LastName:         u.LastName,
			FirstName:        u.FirstName,
			LastNameKana:     u.LastNameKana,
			FirstNameKana:    u.FirstNameKana,
			CreatedAt:        u.CreatedAt,
			UpdatedAt:        u.UpdatedAt,
		}

		users[i] = user
	}

	return &response.UserListResponse{
		Users:  users,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}
