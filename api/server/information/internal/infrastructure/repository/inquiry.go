package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/information/internal/domain/inquiry"
	"gorm.io/gorm/clause"
)

type inquiryRepository struct {
	client *Client
}

//NewInquiryRepository - InquiryRepositoryの生成
func NewInquiryRepository(c *Client) inquiry.Repository {
	return &inquiryRepository{
		client: c,
	}
}

func (r *inquiryRepository) Create(ctx context.Context, i *inquiry.Inquiry) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Omit(clause.Associations).Create(&i).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	return tx.Commit().Error
}
