package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()

	type args struct {
		items  interface{}
		target interface{}
	}
	type want struct {
		isContains bool
		err        error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success []int contain",
			args: args{
				items:  []int{1, 2, 3},
				target: 2,
			},
			want: want{
				isContains: true,
				err:        nil,
			},
		},
		{
			name: "success []int32 contain",
			args: args{
				items:  []int32{1, 2, 3},
				target: int32(2),
			},
			want: want{
				isContains: true,
				err:        nil,
			},
		},
		{
			name: "success []int32 not contain",
			args: args{
				items:  []int32{1, 2, 3},
				target: int32(4),
			},
			want: want{
				isContains: false,
				err:        nil,
			},
		},
		{
			name: "success []int64 contain",
			args: args{
				items:  []int64{1, 2, 3},
				target: int64(2),
			},
			want: want{
				isContains: true,
				err:        nil,
			},
		},
		{
			name: "success []int64 not contain",
			args: args{
				items:  []int64{1, 2, 3},
				target: int64(4),
			},
			want: want{
				isContains: false,
				err:        nil,
			},
		},
		{
			name: "success []string contain",
			args: args{
				items:  []string{"1", "2", "3"},
				target: "2",
			},
			want: want{
				isContains: true,
				err:        nil,
			},
		},
		{
			name: "success []string not contain",
			args: args{
				items:  []string{"1", "2", "3"},
				target: "4",
			},
			want: want{
				isContains: false,
				err:        nil,
			},
		},
		{
			name: "failed other type",
			args: args{
				items:  []int{1, 2, 3},
				target: 2,
			},
			want: want{
				isContains: false,
				err:        ErrUnsupportedType,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Contains(tt.args.items, tt.args.target)
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.isContains, actual)
		})
	}
}
