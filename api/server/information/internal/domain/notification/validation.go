package notification

import "context"

// Validation - Notificationドメインバリデーション
type Validation interface {
	Notification(ctx context.Context, b *Notification) error
}
