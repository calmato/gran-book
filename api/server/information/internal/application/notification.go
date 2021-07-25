package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/application/validation"
	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
)

// NotificationApplication - Notificationアプリケーションのインターフェース
type NotificationApplication interface {
	List(ctx context.Context) ([]*notification.Notification, error)
	Show(ctx context.Context, notificatinID int) (*notification.Notification, error)
	Create(ctx context.Context, in *input.CreateNotification) (*notification.Notification, error)
	Update(ctx context.Context, in *input.UpdateNotification) (*notification.Notification, error)
	Delete(ctx context.Context, notificatinID int) error
}

type notificationApplication struct {
	notificationRequestValidation validation.NotificationRequestValidation
	notifictationService          notification.Service
}

// NewNotificationApplication - NotificationApplicationの生成
func NewNotificationApplication(nrv validation.NotificationRequestValidation,
	ns notification.Service) NotificationApplication {
	return &notificationApplication{
		notificationRequestValidation: nrv,
		notifictationService:          ns,
	}
}

func (a *notificationApplication) List(ctx context.Context) ([]*notification.Notification, error) {
	return a.notifictationService.List(ctx)
}

func (a *notificationApplication) Show(ctx context.Context,
	notificatinID int) (*notification.Notification, error) {
	return a.notifictationService.Show(ctx, notificatinID)
}

func (a *notificationApplication) Create(ctx context.Context,
	in *input.CreateNotification) (*notification.Notification, error) {
	err := a.notificationRequestValidation.CreateNotification(in)
	if err != nil {
		return nil, err
	}

	n := &notification.Notification{
		Title:       in.Title,
		Description: in.Description,
		Importance:  in.Importance,
	}

	err = a.notifictationService.Create(ctx, n)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (a *notificationApplication) Update(ctx context.Context,
	in *input.UpdateNotification) (*notification.Notification, error) {
	err := a.notificationRequestValidation.UpdateNotification(in)
	if err != nil {
		return nil, err
	}

	s, err := a.notifictationService.Show(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	s.Title = in.Title
	s.Description = in.Description
	s.Importance = in.Importance

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (a *notificationApplication) Delete(ctx context.Context,
	notificatinID int) error {
	return a.notifictationService.Delete(ctx, notificatinID)
}
