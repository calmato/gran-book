package validation

import (
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/stretchr/testify/assert"
)

func TestToValidationError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		field   string
		message string
		expect  exception.CustomError
	}{
		{
			name:    "success",
			field:   "id",
			message: exception.RequiredMessage,
			expect: exception.CustomError{
				ErrorCode: exception.InvalidRequestValidation,
				Value:     errInvalidValidation,
				ValidationErrors: []*exception.ValidationError{
					{
						Field:   "id",
						Message: exception.RequiredMessage,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, toValidationError(tt.field, tt.message))
		})
	}
}
