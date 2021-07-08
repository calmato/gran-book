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
	err := r.client.db.Omit(clause.Associations).Create(&i).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
