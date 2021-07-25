package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

type chatDomainValidation struct{}

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

		err := xerrors.New("This message requires either text or image.")
		return exception.InvalidDomainValidation.New(err, ves...)
	}

	return nil
}
