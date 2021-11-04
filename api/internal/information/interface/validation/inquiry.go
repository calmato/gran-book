package validation

import "github.com/calmato/gran-book/api/proto/information"

type inquiryRequestValidation struct{}

func NewInquiryRequestValidation() InquiryRequestValidation {
	return &inquiryRequestValidation{}
}

func (v *inquiryRequestValidation) CreateInquiry(req *information.CreateInquiryRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(information.CreateInquiryRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
