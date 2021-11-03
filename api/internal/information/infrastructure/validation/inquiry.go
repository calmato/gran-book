package inquiry

import (
	"context"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
)

type inquiryDomainValidation struct{}

// NewInquiryDomainVaildation - Inquiry関連のドメインバリデータ
func NewInquiryDomainVaildation() inquiry.Validation {
	return &inquiryDomainValidation{}
}

func (v *inquiryDomainValidation) Inquiry(ctx context.Context, i *inquiry.Inquiry) error {
	return nil
}
