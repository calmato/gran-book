package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
)

type notificationService struct {
	notificationDomainValidation notification.Validation
	notificationRepository       notification.Repository
}

// NewNotificationService - NotificationSeriviceの生成
func NewNotificationService(ndv notification.Validation, nr notification.Repository) notification.Service {
	return &notificationService{
		notificationDomainValidation: ndv,
		notificationRepository:       nr,
	}
}

func (s *notificationService) List(ctx context.Context) ([]*notification.Notification, error) {
	return s.notificationRepository.List(ctx)
}

func (s *notificationService) Show(ctx context.Context, notificatinID int) (*notification.Notification, error) {
	return s.notificationRepository.Show(ctx, notificatinID)
}

func (s *notificationService) Create(ctx context.Context, n *notification.Notification) error {
	currnt := time.Now().Local()

	n.CreatedAt = currnt
	n.UpdatedAt = currnt

	return s.notificationRepository.Create(ctx, n)
}

func (s *notificationService) Update(ctx context.Context, n *notification.Notification) error {
	currnt := time.Now().Local()

	n.UpdatedAt = currnt

	return s.notificationRepository.Update(ctx, n)
}

func (s *notificationService) Delete(ctx context.Context, notificatinID int) error {
	return s.notificationRepository.Delete(ctx, notificatinID)
}
