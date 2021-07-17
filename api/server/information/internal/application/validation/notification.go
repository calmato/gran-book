package validation

import (
	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	"golang.org/x/xerrors"
)

// NotificationRequestValidation - Notifiction関連のリクエストバリデーター
type NotificationRequestValidation interface {
	ShowNotification(in *input.ShowNotification) error
	CreateNotification(in *input.CreateNotification) error
	UpdateNotification(in *input.UpdateNotification) error
	DeleteNotification(in *input.DeleteNotification) error
}

type notificationRequestValidation struct {
	validator RequestValidator
}

// NewNotificationRequestValidation - NotificationRequestValidationの生成
func NewNotificationRequestValidation() NotificationRequestValidation {
	rv := NewRequestValidator()

	return &notificationRequestValidation{
		validator: rv,
	}
}

func (v *notificationRequestValidation) ShowNotification(in *input.ShowNotification) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ShowNotification for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *notificationRequestValidation) CreateNotification(in *input.CreateNotification) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateNotification for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *notificationRequestValidation) UpdateNotification(in *input.UpdateNotification) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateNotification for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *notificationRequestValidation) DeleteNotification(in *input.DeleteNotification) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to DeleteNotification for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
