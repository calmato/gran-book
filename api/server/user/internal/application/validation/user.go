package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// UserRequestValidation - User関連のリクエストバリデータ
type UserRequestValidation interface {
	GetUserProfile(in *input.GetUserProfile) error
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

func (v *userRequestValidation) GetUserProfile(in *input.GetUserProfile) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to GetUserProfile for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
