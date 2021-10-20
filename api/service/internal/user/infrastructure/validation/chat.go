package validation

import (
	"context"
	"errors"

	"github.com/calmato/gran-book/api/service/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/service/pkg/exception"
)

type chatDomainValidation struct{}

var errInvalidChatMessageField = errors.New("validation: this message requires either text or image")

// NewChatDomainValidation - Chat関連のドメインバリデータ
func NewChatDomainValidation() chat.Validation {
	return &chatDomainValidation{}
}

func (v *chatDomainValidation) Room(ctx context.Context, cr *chat.Room) error {
	return nil
}

func (v *chatDomainValidation) Message(ctx context.Context, cm *chat.Message) error {
	if cm.Text == "" && cm.Image == "" {
		ves := []*exception.ValidationError{
			{
				Field:   "text",
				Message: exception.RequiredMessage,
			},
		}

		return exception.ErrInvalidDomainValidation.New(errInvalidChatMessageField, ves...)
	}

	return nil
}
