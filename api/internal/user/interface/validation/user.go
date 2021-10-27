package validation

import (
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
)

type userRequestValidation struct{}

func NewUserRequestValidation() UserRequestValidation {
	return &userRequestValidation{}
}

func (v *userRequestValidation) ListUser(req *user.ListUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.ListUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) ListFollow(req *user.ListFollowRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.ListFollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) ListFollower(req *user.ListFollowerRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.ListFollowerRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) MultiGetUser(req *user.MultiGetUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.MultiGetUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) GetUser(req *user.GetUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.GetUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) GetUserProfile(req *user.GetUserProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(user.GetUserProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) Follow(req *user.FollowRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetUserId() == req.GetFollowerId() {
			return toValidationError("FollowerId", exception.UniqueMessage)
		}

		return nil
	}

	validate := err.(user.FollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) Unfollow(req *user.UnfollowRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetUserId() == req.GetFollowerId() {
			return toValidationError("FollowerId", exception.UniqueMessage)
		}

		return nil
	}

	validate := err.(user.UnfollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
