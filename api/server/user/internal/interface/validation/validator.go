package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
)

var (
	errInvalidValidation  = errors.New("failed to request validation")
	errInternalValidation = errors.New("failed to convert request")
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	return exception.InvalidRequestValidation.New(errInvalidValidation, ve)
}

func toInternalError() error {
	return exception.Unknown.New(errInternalValidation)
}
