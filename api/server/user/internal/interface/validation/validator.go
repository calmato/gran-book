package validation

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"golang.org/x/xerrors"
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	err := xerrors.New("Failed to request validation")
	return exception.InvalidRequestValidation.New(err, ve)
}

func toInternalError() error {
	err := xerrors.New("Failed to convert request")
	return exception.Unknown.New(err)
}
