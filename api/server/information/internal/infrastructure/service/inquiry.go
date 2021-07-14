package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/information/internal/domain/inquiry"
)

type inquiryService struct {
	inquiryRepository inquiry.Repository
}

// NewInquiryService - InquiryServiceの生成
func NewInquiryService(ir inquiry.Repository) inquiry.Service {
	return &inquiryService{
		inquiryRepository: ir,
	}
}

func (s *inquiryService) Create(ctx context.Context, i *inquiry.Inquiry) error {
	current := time.Now().Local()

	i.CreatedAt = current
	i.UpdatedAt = current

	return s.inquiryRepository.Create(ctx, i)
}
