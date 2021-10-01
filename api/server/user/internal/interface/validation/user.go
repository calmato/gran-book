package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
)

type UserRequestValidation interface {
	ListUser(req *pb.ListUserRequest) error
	ListFollow(req *pb.ListFollowRequest) error
	ListFollower(req *pb.ListFollowerRequest) error
	MultiGetUser(req *pb.MultiGetUserRequest) error
	GetUser(req *pb.GetUserRequest) error
	GetUserProfile(req *pb.GetUserProfileRequest) error
	Follow(req *pb.FollowRequest) error
	Unfollow(req *pb.UnfollowRequest) error
}

type userRequestValidation struct{}

func NewUserRequestValidation() UserRequestValidation {
	return &userRequestValidation{}
}

func (v *userRequestValidation) ListUser(req *pb.ListUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) ListFollow(req *pb.ListFollowRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListFollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) ListFollower(req *pb.ListFollowerRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListFollowerRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) MultiGetUser(req *pb.MultiGetUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.MultiGetUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) GetUser(req *pb.GetUserRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetUserRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) GetUserProfile(req *pb.GetUserProfileRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetUserProfileRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) Follow(req *pb.FollowRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetUserId() == req.GetFollowerId() {
			return toValidationError("FollowerId", exception.UniqueMessage)
		}

		return nil
	}

	validate := err.(pb.FollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *userRequestValidation) Unfollow(req *pb.UnfollowRequest) error {
	err := req.Validate()
	if err == nil {
		if req.GetUserId() == req.GetFollowerId() {
			return toValidationError("FollowerId", exception.UniqueMessage)
		}

		return nil
	}

	validate := err.(pb.UnfollowRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
