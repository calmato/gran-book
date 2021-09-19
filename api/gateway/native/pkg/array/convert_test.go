package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		items  interface{}
		expect []string
		err    error
	}{
		{
			name:   "success []int",
			items:  []int{1, 2, 3},
			expect: []string{"1", "2", "3"},
			err:    nil,
		},
		{
			name:   "success []int32",
			items:  []int32{1, 2, 3},
			expect: []string{"1", "2", "3"},
			err:    nil,
		},
		{
			name:   "success []int64",
			items:  []int64{1, 2, 3},
			expect: []string{"1", "2", "3"},
			err:    nil,
		},
		{
			name:   "success []string",
			items:  []string{"1", "2", "3"},
			expect: []string{"1", "2", "3"},
			err:    nil,
		},
		{
			name:   "failed other type",
			items:  []float64{1.0, 2.0, 3.0},
			expect: []string{},
			err:    ErrUnsupportedType,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ConvertStrings(tt.items)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
