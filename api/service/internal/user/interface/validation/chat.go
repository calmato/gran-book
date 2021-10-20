package validation

import "github.com/calmato/gran-book/api/service/proto/chat"

type chatRequestValidation struct{}

func NewChatRequestValidation() ChatRequestValidation {
	return &chatRequestValidation{}
}

func (v *chatRequestValidation) ListRoom(req *chat.ListRoomRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(chat.ListRoomRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) CreateRoom(req *chat.CreateRoomRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(chat.CreateRoomRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) CreateMessage(req *chat.CreateMessageRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(chat.CreateMessageRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *chatRequestValidation) UploadChatImage(req *chat.UploadChatImageRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(chat.UploadChatImageRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
