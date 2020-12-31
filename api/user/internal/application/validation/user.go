package validation

import (
	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// UserRequestValidation - User関連のリクエストバリデータ
type UserRequestValidation interface {
	CreateUser(in *input.CreateUser) error
}

type userRequestValidation struct {
	validator RequestValidator
}

// NewUserRequestValidation - UserRequestValidationの生成
func NewUserRequestValidation() UserRequestValidation {
	rv := NewRequestValidator()

	return &userRequestValidation{
		validator: rv,
	}
}

func (v *userRequestValidation) CreateUser(in *input.CreateUser) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateUser for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
