package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain"
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

func (r *notificationRepository) ListNotification(ctx context.Context, q *domain.ListQuery) ([]*notification.Notification, error) {
	ns := []*notification.Notification{}

	sql := r.client.db.Preload("Notification")
	db := r.client.getListQuery(sql, q)

	err := db.Find(&ns).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return ns, nil
}

func (r *notificationRepository) ShowNotication(ctx context.Context, notificatinID int) (*notification.Notification, error) {
	n := &notification.Notification{}

	err := r.client.db.Preload("Notification").First(n, "id = ?", notificatinID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return n, nil
}

func (r *notificationRepository) CreateNotification(ctx context.Context, n *notification.Notification) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Omit(clause.Associations).Create(&n).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	return tx.Commit().Error
}

func (r *notificationRepository) UpdateNotification(ctx context.Context, n *notification.Notification) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Omit(clause.Associations).Save(&b).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	return tx.Commit().Error
}

func (r *notificationRepository) DeleteNotification(ctx context.Context, notificatinID int) error {
	err := r.client.db.Where("id = ?", notificatinID).Delete(&notification.Notification{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
