package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// ChatRequestValidation - Chat関連のリクエストバリデータ
type ChatRequestValidation interface {
	CreateRoom(in *input.CreateRoom) error
}

type chatRequestValidation struct {
	validator RequestValidator
}

// NewChatRequestValidation - ChatRequestValidationの生成
func NewChatRequestValidation() ChatRequestValidation {
	rv := NewRequestValidator()

	return &chatRequestValidation{
		validator: rv,
	}
}

func (v *chatRequestValidation) CreateRoom(in *input.CreateRoom) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateRoom for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
