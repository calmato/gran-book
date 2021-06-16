package validation

import (
	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	"golang.org/x/xerrors"
)

// InquiryRequestValidation - Inquiry関連のリクエストバリデータ
type InquiryRequestValidation interface {
	CreateInquiry(in *input.CreateInquiry) error
}

type inquiryRequestValidation struct {
	validator RequestValidator
}

// NewInquiryRequestValidation - InquiryRequestValidationの生成
func NewInquiryRequestValidation() InquiryRequestValidation {
	rv := NewRequestValidator()

	return &inquiryRequestValidation{
		validator: rv,
	}
}

func (v *inquiryRequestValidation) CreateInquiry(in *input.CreateInquiry) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to Inquiry for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
