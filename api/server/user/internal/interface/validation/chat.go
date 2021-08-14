package validation

import (
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

type ChatRequestValidation interface {
	ListChatRoom(req *pb.ListChatRoomRequest) error
	CreateChatRoom(req *pb.CreateChatRoomRequest) error
	CreateChatMessage(req *pb.CreateChatMessageRequest) error
	UploadChatImage(req *pb.UploadChatImageRequest) error
}

type chatRequestValidation struct{}

func NewChatRequestValidation() ChatRequestValidation {
	return &chatRequestValidation{}
}

func (v *chatRequestValidation) ListChatRoom(req *pb.ListChatRoomRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ListChatRoomRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *chatRequestValidation) CreateChatRoom(req *pb.CreateChatRoomRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateChatRoomRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *chatRequestValidation) CreateChatMessage(req *pb.CreateChatMessageRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateChatMessageRequestValidationError); ok {
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
