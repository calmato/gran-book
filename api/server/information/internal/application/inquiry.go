package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/application/validation"
	"github.com/calmato/gran-book/api/server/information/internal/domain/inquiry"
)

// InquiryApplication - Inquiryアプリケーションのインターフェース
type InquiryApplication interface {
	Create(ctx context.Context, in *input.CreateInquiry) (*inquiry.Inquiry, error)
}

type inquiryApplication struct {
	inquiryRequestValidation validation.InquiryRequestValidation
	inquiryService           inquiry.Service
}

// NewInquiryApplication - InquiryApplicationの生成
func NewInquiryApplication(irv validation.InquiryRequestValidation, is inquiry.Service) InquiryApplication {
	return &inquiryApplication{
		inquiryRequestValidation: irv,
		inquiryService:           is,
	}
}

func (ia *inquiryApplication) Create(ctx context.Context, in *input.CreateInquiry) (*inquiry.Inquiry, error) {
	err := ia.inquiryRequestValidation.CreateInquiry(in)
	if err != nil {
		return nil, err
	}

	i := &inquiry.Inquiry{
		SenderID:    in.SenderID,
		Subject:     in.Subject,
		Description: in.Description,
		Email:       in.Email,
	}

	err = ia.inquiryService.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
