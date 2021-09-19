package entity

import (
	"fmt"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomError(t *testing.T) {
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
			require.IsType(t, CustomError{}, actual)
		})
	}
}

func TestCustomError_Error(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		code   ErrorCode
		err    error
		expect error
	}{
		{
			name:   "bad request",
			code:   ErrBadRequest,
			err:    test.ErrMock,
			expect: test.ErrMock,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.code.New(tt.err)
			require.Error(t, actual)

			err, ok := actual.(CustomError)
			require.True(t, ok, fmt.Sprintln("actual is not CustomError type"))

			assert.Equal(t, tt.expect.Error(), err.Error())
		})
	}
}

func TestCustomError_Code(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		code   ErrorCode
		err    error
		expect ErrorCode
	}{
		{
			name:   "bad request",
			code:   ErrBadRequest,
			err:    test.ErrMock,
			expect: ErrBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.code.New(tt.err)
			require.Error(t, actual)

			err, ok := actual.(CustomError)
			require.True(t, ok, fmt.Sprintln("actual is not CustomError type"))

			assert.Equal(t, tt.expect, err.Code())
		})
	}
}
