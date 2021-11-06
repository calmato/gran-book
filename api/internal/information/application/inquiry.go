package application

import (
	"context"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
)

type inquiryApplication struct {
	inquiryDomainValidation inquiry.Validation
	inquiryRepository       inquiry.Repository
}

// NewInquiryApplication - InquiryApplicationの生成
func NewInquiryApplication(idv inquiry.Validation, ir inquiry.Repository) InquiryApplication {
	return &inquiryApplication{
		inquiryDomainValidation: idv,
		inquiryRepository:       ir,
	}
}

func (a *inquiryApplication) Create(ctx context.Context, i *inquiry.Inquiry) error {
	err := a.inquiryDomainValidation.Inquiry(ctx, i)
	if err != nil {
		return err
	}

	return a.inquiryRepository.Create(ctx, i)
}
