package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/domain/inquiry"
)

// InquiryApplication - Inquiryアプリケーションのインターフェース
type InquiryApplication interface {
	Create(ctx context.Context, in *input.CreateInquiry) (*inquiry.Inquiry, error)
}

type inquiryApplication struct {
	inquiryService inquiry.Service
}

// NewInquiryApplication - InquiryApplicationの生成
func NewInquiryApplication(is inquiry.Service) InquiryApplication {
	return &inquiryApplication{
		inquiryService: is,
	}
}

func (ia *inquiryApplication) Create(ctx context.Context, in *input.CreateInquiry) (*inquiry.Inquiry, error) {
	i := &inquiry.Inquiry{
		SenderID:    in.SenderID,
		Subject:     in.Subject,
		Description: in.Description,
		Email:       in.Email,
	}

	err := ia.inquiryService.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
