package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
	"gorm.io/gorm/clause"
)

type notificationRepository struct {
	client *Client
}

// NewNotificationRepository - NotificationRepositoryの生成
func NewNotificationRepository(c *Client) notification.Repository {
	return &notificationRepository{
		client: c,
	}
}

func (r *notificationRepository) List(ctx context.Context) ([]*notification.Notification, error) {
	ns := []*notification.Notification{}

	err := r.client.db.Find(&ns).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return ns, nil
}

func (r *notificationRepository) Show(ctx context.Context, notificatinID int) (*notification.Notification, error) {
	n := &notification.Notification{}

	err := r.client.db.First(n, "id = ?", notificatinID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return n, nil
}

func (r *notificationRepository) Create(ctx context.Context, n *notification.Notification) error {
	err := r.client.db.Omit(clause.Associations).Create(&n).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *notificationRepository) Update(ctx context.Context, n *notification.Notification) error {
	err := r.client.db.Omit(clause.Associations).Save(&r).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *notificationRepository) Delete(ctx context.Context, notificatinID int) error {
	err := r.client.db.Where("id = ?", notificatinID).Delete(&notification.Notification{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
