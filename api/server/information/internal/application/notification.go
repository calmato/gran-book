package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/application/validation"
	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
)

// NotificationApplication - Notificationアプリケーションのインターフェース
type NotificationApplication interface {
	Create(ctx context.Context, in *input.CreateNotification) error
}

type notificationApplication struct {
	notificationRequestValidation validation.NotificationRequestValidation
	notifictationService          notification.Service
}

// NewNotificationApplication - NotificationApplicationの生成
func NewNotificationApplication(nrv validation.NotificationRequestValidation, ns notification.Service) NotificationApplication {
	return &notificationApplication{
		notificationRequestValidation: nrv,
		notifictationService:          ns,
	}
}

func (n *notificationApplication) Create(ctx context.Context, in *input.CreateNotification) error {
	err := n.notificationRequestValidation.CreateNotification(in)
	if err != nil {
		return err
	}

	u := &notification.Notification{
		Title:       in.Title,
		Description: in.Description,
		Importance:  in.Importance,
	}

	// TODO: valication check

	err = n.notifictationService.Create(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
