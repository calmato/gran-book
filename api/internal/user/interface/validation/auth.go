package validation

import (
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
)

type authRequestValidation struct{}

func NewAuthRequestValidation() AuthRequestValidation {
	return &authRequestValidation{}
}

func (v *authRequestValidation) CreateAuth(req *user.CreateAuthRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(user.CreateAuthRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthEmail(req *user.UpdateAuthEmailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UpdateAuthEmailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthPassword(req *user.UpdateAuthPasswordRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(user.UpdateAuthPasswordRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthProfile(req *user.UpdateAuthProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UpdateAuthProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthAddress(req *user.UpdateAuthAddressRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UpdateAuthAddressRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UploadAuthThumbnail(req *user.UploadAuthThumbnailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.UploadAuthThumbnailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) RegisterAuthDevice(req *user.RegisterAuthDeviceRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.RegisterAuthDeviceRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
