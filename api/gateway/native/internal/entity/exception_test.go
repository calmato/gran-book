package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorCode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		code   ErrorCode
		err    error
		expect error
	}{
		{
			name: "bad request",
			code: ErrBadRequest,
			err:  test.ErrMock,
			expect: CustomError{
				code:  ErrBadRequest,
				value: test.ErrMock,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.code.New(tt.err)
			require.Error(t, actual)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.err.Error(), actual.Error())
		})
	}
}
