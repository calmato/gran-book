//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/interface/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/calmato/gran-book/api/proto/user"
)

var (
	errInvalidValidation = errors.New("validation: failed to convert request")
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	return exception.ErrInvalidRequestValidation.New(errInvalidValidation, ve)
}

type AuthRequestValidation interface {
	CreateAuth(req *user.CreateAuthRequest) error
	UpdateAuthEmail(req *user.UpdateAuthEmailRequest) error
	UpdateAuthPassword(req *user.UpdateAuthPasswordRequest) error
	UpdateAuthProfile(req *user.UpdateAuthProfileRequest) error
	UpdateAuthAddress(req *user.UpdateAuthAddressRequest) error
	UploadAuthThumbnail(req *user.UploadAuthThumbnailRequest) error
	RegisterAuthDevice(req *user.RegisterAuthDeviceRequest) error
}

type UserRequestValidation interface {
	ListUser(req *user.ListUserRequest) error
	ListFollow(req *user.ListFollowRequest) error
	ListFollower(req *user.ListFollowerRequest) error
	MultiGetUser(req *user.MultiGetUserRequest) error
	GetUser(req *user.GetUserRequest) error
	GetUserProfile(req *user.GetUserProfileRequest) error
	Follow(req *user.FollowRequest) error
	Unfollow(req *user.UnfollowRequest) error
}

type AdminRequestValidation interface {
	ListAdmin(req *user.ListAdminRequest) error
	GetAdmin(req *user.GetAdminRequest) error
	CreateAdmin(req *user.CreateAdminRequest) error
	UpdateAdminContact(req *user.UpdateAdminContactRequest) error
	UpdateAdminPassword(req *user.UpdateAdminPasswordRequest) error
	UpdateAdminProfile(req *user.UpdateAdminProfileRequest) error
	UploadAdminThumbnail(req *user.UploadAdminThumbnailRequest) error
	DeleteAdmin(req *user.DeleteAdminRequest) error
}

type ChatRequestValidation interface {
	ListRoom(req *chat.ListRoomRequest) error
	CreateRoom(req *chat.CreateRoomRequest) error
	CreateMessage(req *chat.CreateMessageRequest) error
	UploadChatImage(req *chat.UploadChatImageRequest) error
}
