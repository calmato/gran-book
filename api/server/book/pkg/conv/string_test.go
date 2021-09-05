package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCamelToSnake(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title     string
		input     string
		expect    string
		expectErr error
	}{
		{
			title:     "success: RememberMe",
			input:     "RememberMe",
			expect:    "remember_me",
			expectErr: nil,
		},
		{
			title:     "success: HTTPRequset",
			input:     "HTTPRequest",
			expect:    "h_t_t_p_request",
			expectErr: nil,
		},
		{
			title:     "success: gran-book",
			input:     "gran-book",
			expect:    "gran-book",
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			got, err := CamelToSnake(tt.input)
			if err != nil {
				require.Error(t, err)
				assert.Errorf(t, tt.expectErr, ErrStringConvert.Error())
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, got)
		})
	}
}
