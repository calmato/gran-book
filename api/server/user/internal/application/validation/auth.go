package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// AuthRequestValidation - Auth関連のリクエストバリデータ
type AuthRequestValidation interface {
	CreateAuth(in *input.CreateAuth) error
	UpdateAuthEmail(in *input.UpdateAuthEmail) error
	UpdateAuthPassword(in *input.UpdateAuthPassword) error
	UpdateAuthProfile(in *input.UpdateAuthProfile) error
	UpdateAuthAddress(in *input.UpdateAuthAddress) error
}

type authRequestValidation struct {
	validator RequestValidator
}

// NewAuthRequestValidation - AuthRequestValidationの生成
func NewAuthRequestValidation() AuthRequestValidation {
	rv := NewRequestValidator()

	return &authRequestValidation{
		validator: rv,
	}
}

func (v *authRequestValidation) CreateAuth(in *input.CreateAuth) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateAuth for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *authRequestValidation) UpdateAuthEmail(in *input.UpdateAuthEmail) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAuthEmail for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *authRequestValidation) UpdateAuthPassword(in *input.UpdateAuthPassword) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAuthPassword for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *authRequestValidation) UpdateAuthProfile(in *input.UpdateAuthProfile) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAuthProfile for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *authRequestValidation) UpdateAuthAddress(in *input.UpdateAuthAddress) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAuthAddress for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
