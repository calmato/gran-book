package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/user/proto/user"
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
	if err != nil {
		if err, ok := err.(pb.CreateAuthRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	if req.GetPassword() != req.GetPasswordConfirmation() {
		return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
	}

	return nil
}

func (v *authRequestValidation) UpdateAuthEmail(req *pb.UpdateAuthEmailRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAuthEmailRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *authRequestValidation) UpdateAuthPassword(req *pb.UpdateAuthPasswordRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAuthPasswordRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	if req.GetPassword() != req.GetPasswordConfirmation() {
		return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
	}

	return nil
}

func (v *authRequestValidation) UpdateAuthProfile(req *pb.UpdateAuthProfileRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAuthProfileRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *authRequestValidation) UpdateAuthAddress(req *pb.UpdateAuthAddressRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAuthAddressRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *authRequestValidation) UploadAuthThumbnail(req *pb.UploadAuthThumbnailRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UploadAuthThumbnailRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *authRequestValidation) RegisterAuthDevice(req *pb.RegisterAuthDeviceRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.RegisterAuthDeviceRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}
