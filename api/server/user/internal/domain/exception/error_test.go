package exception

import (
	"testing"

	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		code   ErrorCode
		err    error
		ves    []*ValidationError
		expect CustomError
	}{
		{
			name: "success",
			code: ErrorInDatastore,
			err:  test.ErrMock,
			ves:  []*ValidationError{},
			expect: CustomError{
				ErrorCode:        ErrorInDatastore,
				Value:            test.ErrMock,
				ValidationErrors: []*ValidationError{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.code.New(tt.err, tt.ves...)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestCustomError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    CustomError
		expect string
	}{
		{
			name: "success",
			err: CustomError{
				ErrorCode:        ErrorInDatastore,
				Value:            test.ErrMock,
				ValidationErrors: []*ValidationError{},
			},
			expect: test.ErrMock.Error(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.err.Error())
		})
	}
}

func TestCustomError_Code(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    CustomError
		expect ErrorCode
	}{
		{
			name: "success",
			err: CustomError{
				ErrorCode:        ErrorInDatastore,
				Value:            test.ErrMock,
				ValidationErrors: []*ValidationError{},
			},
			expect: ErrorInDatastore,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.err.Code())
		})
	}
}

func TestCustomError_Validations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    CustomError
		expect []*ValidationError
	}{
		{
			name: "success",
			err: CustomError{
				ErrorCode:        ErrorInDatastore,
				Value:            test.ErrMock,
				ValidationErrors: []*ValidationError{},
			},
			expect: []*ValidationError{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.err.Validations())
		})
	}
}
