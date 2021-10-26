package validation

import (
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
)

type adminRequestValidation struct{}

func NewAdminRequestValidation() AdminRequestValidation {
	return &adminRequestValidation{}
}

func (v *adminRequestValidation) ListAdmin(req *user.ListAdminRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.ListAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) GetAdmin(req *user.GetAdminRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.GetAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) CreateAdmin(req *user.CreateAdminRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(user.CreateAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminContact(req *user.UpdateAdminContactRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UpdateAdminContactRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminPassword(req *user.UpdateAdminPasswordRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(user.UpdateAdminPasswordRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminProfile(req *user.UpdateAdminProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UpdateAdminProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UploadAdminThumbnail(req *user.UploadAdminThumbnailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UploadAdminThumbnailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) DeleteAdmin(req *user.DeleteAdminRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.DeleteAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
