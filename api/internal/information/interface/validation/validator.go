//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/information/interface/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/information"
)

var (
	errInvalidValidation = errors.New("validation: failed to invalid validation")
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	return exception.ErrInvalidRequestValidation.New(errInvalidValidation, ve)
}

type InquiryRequestValidation interface {
	CreateInquiry(req *information.CreateInquiryRequest) error
}
