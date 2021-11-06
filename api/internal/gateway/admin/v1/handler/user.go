package handler

import (
	"context"
	"net/http"

	"github.com/calmato/gran-book/api/internal/gateway/admin/v1/entity"
	response "github.com/calmato/gran-book/api/internal/gateway/admin/v1/response"
	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/conv"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) listUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	field := ctx.DefaultQuery("field", "")
	value := ctx.DefaultQuery("value", "")
	by := ctx.DefaultQuery("by", "")
	direction := ctx.DefaultQuery("direction", "")

	us, total, err := h.userListUser(c, limit, offset, field, value, by, direction)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.UserListResponse{
		Users:  entity.NewUsers(us),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) getUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	u, err := h.userGetUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.UserResponse{
		User: entity.NewUser(u),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) userListUser(
	ctx context.Context,
	limit, offset int64,
	field, value, by, direction string,
) (gentity.Users, int64, error) {
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
		orderBy, err := conv.CamelToSnake(by)
		if err != nil {
			err := exception.ErrInvalidArgument.New(err)
			return nil, 0, err
		}

		order := &user.Order{
			Field:   orderBy,
			OrderBy: gentity.NewOrderByByValue(direction).Proto(),
		}
		in.Order = order
	}

	out, err := h.User.ListUser(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewUsers(out.Users), out.Total, nil
}

func (h *apiV1Handler) userGetUser(ctx context.Context, userID string) (*gentity.User, error) {
	in := &user.GetUserRequest{
		UserId: userID,
	}

	out, err := h.User.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewUser(out.User), nil
}
