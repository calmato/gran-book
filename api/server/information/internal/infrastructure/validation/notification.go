package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
)

type notificationDomainValidation struct{}

func NewNotificationDomainValidation() notification.Validation {
	return &notificationDomainValidation{}
}

func (dv *notificationDomainValidation) Notification(ctx context.Context, n *notification.Notification) error {
	return nil
}
