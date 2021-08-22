package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	err := errors.New("failed to request validation")
	return exception.InvalidRequestValidation.New(err, ve)
}

func toInternalError() error {
	err := errors.New("failed to convert request")
	return exception.Unknown.New(err)
}
