package notification

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain"
)

// Service - Notifiationサービス
type Service interface {
	ListNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	ShowNotication(ctx context.Context, notificatinID int) (*Notification, error)
	CreateNotification(ctx context.Context, n *Notification) error
	UpdateNotification(ctx context.Context, n *Notification) error
	DeleteNotification(ctx context.Context, notificatinID int) error
}
