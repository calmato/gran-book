package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

type AdminRequestValidation interface {
	ListAdmin(req *pb.ListAdminRequest) error
	GetAdmin(req *pb.GetAdminRequest) error
	CreateAdmin(req *pb.CreateAdminRequest) error
	UpdateAdminContact(req *pb.UpdateAdminContactRequest) error
	UpdateAdminPassword(req *pb.UpdateAdminPasswordRequest) error
	UpdateAdminProfile(req *pb.UpdateAdminProfileRequest) error
	UploadAdminThumbnail(req *pb.UploadAdminThumbnailRequest) error
	DeleteAdmin(req *pb.DeleteAdminRequest) error
}

type adminRequestValidation struct{}

func NewAdminRequestValidation() AdminRequestValidation {
	return &adminRequestValidation{}
}

func (v *adminRequestValidation) ListAdmin(req *pb.ListAdminRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ListAdminRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *adminRequestValidation) GetAdmin(req *pb.GetAdminRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.GetAdminRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *adminRequestValidation) CreateAdmin(req *pb.CreateAdminRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateAdminRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	if req.GetPassword() != req.GetPasswordConfirmation() {
		return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
	}

	return nil
}

func (v *adminRequestValidation) UpdateAdminContact(req *pb.UpdateAdminContactRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAdminContactRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *adminRequestValidation) UpdateAdminPassword(req *pb.UpdateAdminPasswordRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAdminPasswordRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	if req.GetPassword() != req.GetPasswordConfirmation() {
		return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
	}

	return nil
}

func (v *adminRequestValidation) UpdateAdminProfile(req *pb.UpdateAdminProfileRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateAdminProfileRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *adminRequestValidation) UploadAdminThumbnail(req *pb.UploadAdminThumbnailRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UploadAdminThumbnailRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *adminRequestValidation) DeleteAdmin(req *pb.DeleteAdminRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.DeleteAdminRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}
