package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
)

type chatDomainValidation struct{}

// NewChatDomainValidation - Chat関連のドメインバリデータ
func NewChatDomainValidation() chat.Validation {
	return &chatDomainValidation{}
}

func (v *chatDomainValidation) Room(ctx context.Context, cr *chat.Room) error {
	return nil
}
