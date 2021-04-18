package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// UserRequestValidation - User関連のリクエストバリデータ
type UserRequestValidation interface {
	ListUser(in *input.ListUser) error
	ListFollow(in *input.ListFollow) error
	ListFollower(in *input.ListFollower) error
	SearchUser(in *input.SearchUser) error
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

func (v *userRequestValidation) ListUser(in *input.ListUser) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListUser for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *userRequestValidation) ListFollow(in *input.ListFollow) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListFollow for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *userRequestValidation) ListFollower(in *input.ListFollower) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListFollower for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *userRequestValidation) SearchUser(in *input.SearchUser) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to SearchUser for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
