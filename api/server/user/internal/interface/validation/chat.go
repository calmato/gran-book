package validation

import (
	pb "github.com/calmato/gran-book/api/server/user/proto/service/chat"
)

type ChatRequestValidation interface {
	ListRoom(req *pb.ListRoomRequest) error
	CreateRoom(req *pb.CreateRoomRequest) error
	CreateMessage(req *pb.CreateMessageRequest) error
	UploadChatImage(req *pb.UploadChatImageRequest) error
}

type chatRequestValidation struct{}

func NewChatRequestValidation() ChatRequestValidation {
	return &chatRequestValidation{}
}

func (v *chatRequestValidation) ListRoom(req *pb.ListRoomRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListRoomRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) CreateRoom(req *pb.CreateRoomRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.CreateRoomRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) CreateMessage(req *pb.CreateMessageRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.CreateMessageRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) UploadChatImage(req *pb.UploadChatImageRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UploadChatImageRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
