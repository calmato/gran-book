package notification

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain"
)

// Repository - Notificationリポジトリ
type Repository interface {
	ListNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	ShowNotication(ctx context.Context, notificatinID int) (*Notification, error)
	CreateNotification(ctx context.Context, n *domain.ListQuery) error
	UpdateNotification(ctx context.Context, n *domain.ListQuery) error
	DeleteNotification(ctx context.Context, notificatinID int) error
}
