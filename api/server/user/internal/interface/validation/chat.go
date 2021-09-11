package validation

import (
	pb "github.com/calmato/gran-book/api/server/user/proto/chat"
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
	if err != nil {
		if err, ok := err.(pb.ListRoomRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *chatRequestValidation) CreateRoom(req *pb.CreateRoomRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateRoomRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *chatRequestValidation) CreateMessage(req *pb.CreateMessageRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateMessageRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *chatRequestValidation) UploadChatImage(req *pb.UploadChatImageRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UploadChatImageRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}
