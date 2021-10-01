package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
)

type AuthRequestValidation interface {
	CreateAuth(req *pb.CreateAuthRequest) error
	UpdateAuthEmail(req *pb.UpdateAuthEmailRequest) error
	UpdateAuthPassword(req *pb.UpdateAuthPasswordRequest) error
	UpdateAuthProfile(req *pb.UpdateAuthProfileRequest) error
	UpdateAuthAddress(req *pb.UpdateAuthAddressRequest) error
	UploadAuthThumbnail(req *pb.UploadAuthThumbnailRequest) error
	RegisterAuthDevice(req *pb.RegisterAuthDeviceRequest) error
}

type authRequestValidation struct{}

func NewAuthRequestValidation() AuthRequestValidation {
	return &authRequestValidation{}
}

func (v *authRequestValidation) CreateAuth(req *pb.CreateAuthRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(pb.CreateAuthRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthEmail(req *pb.UpdateAuthEmailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateAuthEmailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthPassword(req *pb.UpdateAuthPasswordRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(pb.UpdateAuthPasswordRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthProfile(req *pb.UpdateAuthProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateAuthProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UpdateAuthAddress(req *pb.UpdateAuthAddressRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateAuthAddressRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) UploadAuthThumbnail(req *pb.UploadAuthThumbnailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UploadAuthThumbnailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *authRequestValidation) RegisterAuthDevice(req *pb.RegisterAuthDeviceRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.RegisterAuthDeviceRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
