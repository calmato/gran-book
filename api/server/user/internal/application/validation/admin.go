package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

// AdminRequestValidation - Admin関連のリクエストバリデータ
type AdminRequestValidation interface {
	ListAdmin(in *input.ListAdmin) error
	SearchAdmin(in *input.SearchAdmin) error
	CreateAdmin(in *input.CreateAdmin) error
	UpdateAdminContact(in *input.UpdateAdminContact) error
	UpdateAdminPassword(in *input.UpdateAdminPassword) error
	UpdateAdminProfile(in *input.UpdateAdminProfile) error
	UploadAdminThumbnail(in *input.UploadAdminThumbnail) error
}

type adminRequestValidation struct {
	validator RequestValidator
}

// NewAdminRequestValidation - AdminRequestValidationの生成
func NewAdminRequestValidation() AdminRequestValidation {
	rv := NewRequestValidator()

	return &adminRequestValidation{
		validator: rv,
	}
}

func (v *adminRequestValidation) ListAdmin(in *input.ListAdmin) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListAdmin for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) SearchAdmin(in *input.SearchAdmin) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to SearchAdmin for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) CreateAdmin(in *input.CreateAdmin) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateAdmin for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) UpdateAdminContact(in *input.UpdateAdminContact) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAdminContact for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) UpdateAdminPassword(in *input.UpdateAdminPassword) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAdminPassword for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) UpdateAdminProfile(in *input.UpdateAdminProfile) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateAdminProfile for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *adminRequestValidation) UploadAdminThumbnail(in *input.UploadAdminThumbnail) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UploadAdminThumbnail for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
