package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
)

var (
	errInvalidValidation = errors.New("validation: failed to convert request")
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	return exception.InvalidRequestValidation.New(errInvalidValidation, ve)
}
