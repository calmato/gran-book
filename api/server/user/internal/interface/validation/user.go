package validation

import (
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

type UserRequestValidation interface {
	ListUser(req *pb.ListUserRequest) error
	ListFollow(req *pb.ListFollowRequest) error
	ListFollower(req *pb.ListFollowerRequest) error
	MultiGetUser(req *pb.MultiGetUserRequest) error
	GetUser(req *pb.GetUserRequest) error
	GetUserProfile(req *pb.GetUserProfileRequest) error
	RegisterFollow(req *pb.RegisterFollowRequest) error
	UnregisterFollow(req *pb.UnregisterFollowRequest) error
}

type userRequestValidation struct{}

func NewUserRequestValidation() UserRequestValidation {
	return &userRequestValidation{}
}

func (v *userRequestValidation) ListUser(req *pb.ListUserRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.ListUserRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) ListFollow(req *pb.ListFollowRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.ListFollowRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) ListFollower(req *pb.ListFollowerRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.ListFollowerRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) MultiGetUser(req *pb.MultiGetUserRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.MultiGetUserRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) GetUser(req *pb.GetUserRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.GetUserRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) GetUserProfile(req *pb.GetUserProfileRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.GetUserProfileRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) RegisterFollow(req *pb.RegisterFollowRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.RegisterFollowRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}

func (v *userRequestValidation) UnregisterFollow(req *pb.UnregisterFollowRequest) error {
	err := req.Validate()
	if err != nil {
		return nil
	}

	if err, ok := err.(pb.UnregisterFollowRequestValidationError); ok {
		return toValidationError(err.Field(), err.Reason())
	}

	return toInternalError()
}
