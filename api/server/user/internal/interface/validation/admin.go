package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
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
	if err == nil {
		return nil
	}

	validate := err.(pb.ListAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) GetAdmin(req *pb.GetAdminRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) CreateAdmin(req *pb.CreateAdminRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(pb.CreateAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminContact(req *pb.UpdateAdminContactRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateAdminContactRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminPassword(req *pb.UpdateAdminPasswordRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetPassword() != req.GetPasswordConfirmation() {
			return toValidationError("PasswordConfirmation", exception.EqFieldMessage)
		}

		return nil
	}

	validate := err.(pb.UpdateAdminPasswordRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UpdateAdminProfile(req *pb.UpdateAdminProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateAdminProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) UploadAdminThumbnail(req *pb.UploadAdminThumbnailRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UploadAdminThumbnailRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *adminRequestValidation) DeleteAdmin(req *pb.DeleteAdminRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.DeleteAdminRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
