package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		expectLocation *time.Location
	}{
		{
			name:           "success",
			expectLocation: jst,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectLocation, Now().Location())
		})
	}
}

func TestBeggingOfMonth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  time.Time
		expect time.Time
	}{
		{
			name:   "success",
			input:  time.Date(2021, time.Month(8), 15, 12, 0, 0, 0, jst),
			expect: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, jst),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, BeginningOfMonth(tt.input))
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  time.Time
		expect time.Time
	}{
		{
			name:   "success",
			input:  time.Date(2021, time.Month(8), 15, 12, 0, 0, 0, jst),
			expect: time.Date(2021, time.Month(8), 31, 23, 59, 59, 0, jst),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, EndOfMonth(tt.input))
		})
	}
}

func TestParseTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		input     string
		expect    time.Time
		expectErr bool
	}{
		{
			name:      "success",
			input:     "2021-08-15 01:30:00",
			expect:    time.Date(2021, time.Month(8), 15, 1, 30, 0, 0, jst),
			expectErr: false,
		},
		{
			name:      "failed to invalid format",
			input:     "20210815 013000",
			expect:    time.Time{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ParseTime(tt.input)
			require.Equal(t, tt.expectErr, err != nil)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestParseDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		input     string
		expect    time.Time
		expectErr bool
	}{
		{
			name:      "success",
			input:     "2021-08-15",
			expect:    time.Date(2021, time.Month(8), 15, 0, 0, 0, 0, jst),
			expectErr: false,
		},
		{
			name:      "success",
			input:     "20210815",
			expect:    time.Time{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ParseDate(tt.input)
			require.Equal(t, tt.expectErr, err != nil)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFormatTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  time.Time
		expect string
	}{
		{
			name:   "success",
			input:  time.Date(2021, time.Month(8), 15, 1, 30, 0, 0, jst),
			expect: "2021-08-15 01:30:00",
		},
		{
			name:   "success is zero value",
			input:  time.Time{},
			expect: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, FormatTime(tt.input))
		})
	}
}

func TestFormatDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  time.Time
		expect string
	}{
		{
			name:   "success",
			input:  time.Date(2021, time.Month(8), 15, 1, 30, 0, 0, jst),
			expect: "2021-08-15",
		},
		{
			name:   "success is zero value",
			input:  time.Time{},
			expect: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, FormatDate(tt.input))
		})
	}
}
