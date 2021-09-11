package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  time.Time
		expect *Time
	}{
		{
			name:  "success",
			input: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, jst),
			expect: &Time{
				time: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, jst),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, New(tt.input))
		})
	}
}

func TestTime_BeggingOfMonth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		time   *Time
		expect time.Time
	}{
		{
			name: "success",
			time: &Time{
				time: time.Date(2021, time.Month(8), 15, 12, 0, 0, 0, jst),
			},
			expect: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, jst),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.time.BeginningOfMonth())
		})
	}
}

func TestTime_EndOfMonth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		time   *Time
		expect time.Time
	}{
		{
			name: "success",
			time: &Time{
				time: time.Date(2021, time.Month(8), 15, 12, 0, 0, 0, jst),
			},
			expect: time.Date(2021, time.Month(8), 31, 23, 59, 59, 0, jst),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.time.EndOfMonth())
		})
	}
}
