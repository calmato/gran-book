package repository

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/exception"
	"gorm.io/gorm"
)

type inquiryRepository struct {
	client *database.Client
	now    func() time.Time
}

// NewInquiryRepository - InquiryRepositoryの生成
func NewInquiryRepository(c *database.Client, now func() time.Time) inquiry.Repository {
	return &inquiryRepository{
		client: c,
		now:    now,
	}
}

func (r *inquiryRepository) Create(ctx context.Context, i *inquiry.Inquiry) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.create(tx, i)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *inquiryRepository) create(tx *gorm.DB, i *inquiry.Inquiry) error {
	now := r.now()
	i.CreatedAt = now
	i.UpdatedAt = now

	return tx.Table(inquiryTable).Create(&i).Error
}
