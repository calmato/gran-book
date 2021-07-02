package notification

import (
	"context"
)

// Repository - Notificationリポジトリ
type Repository interface {
	List(ctx context.Context) ([]*Notification, error)
	Show(ctx context.Context, notificatinID int) (*Notification, error)
	Create(ctx context.Context, n *Notification) error
	Update(ctx context.Context, n *Notification) error
	Delete(ctx context.Context, notificatinID int) error
}
