package notification

import (
	"context"
)

// Repository - Notificationリポジトリ
type Repository interface {
	CreateNotification(ctx context.Context, n *Notification) error
}
