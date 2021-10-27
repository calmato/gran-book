package v2

import (
	"context"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/proto/user"
)

func (h *apiV2Handler) userMultiGetUser(ctx context.Context, userIDs []string) (gentity.Users, error) {
	in := &user.MultiGetUserRequest{
		UserIds: userIDs,
	}
	out, err := h.User.MultiGetUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewUsers(out.Users), nil
}
