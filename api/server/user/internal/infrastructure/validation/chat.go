package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
)

type chatDomainValidation struct{}

func (v *chatDomainValidation) Room(ctx context.Context, cr *chat.Room) error {
	return nil
}
