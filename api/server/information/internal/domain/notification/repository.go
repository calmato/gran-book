package notification

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain"
)

// Repository - Informationリポジトリ
type Repository interface {
	ListNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	ShowNotication(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	CreateNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	UpdateNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
	DeleteNotification(ctx context.Context, q *domain.ListQuery) ([]*Notification, error)
}
