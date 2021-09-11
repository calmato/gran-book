package datetime

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToString(t *testing.T) {
	testCases := map[string]struct {
		Input    time.Time
		Expected string
	}{
		"ok": {
			Input:    time.Date(2020, 1, 1, 1, 1, 1, 0, time.Local),
			Expected: "2020-01-01 01:01:01",
		},
		"ok_is_zero": {
			Input:    time.Time{},
			Expected: "",
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			got := TimeToString(tc.Input)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	testCases := map[string]struct {
		Input    time.Time
		Expected string
	}{
		"ok": {
			Input:    time.Date(2020, 1, 1, 1, 1, 1, 0, time.Local),
			Expected: "2020-01-01",
		},
		"ok_is_zero": {
			Input:    time.Time{},
			Expected: "",
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			got := DateToString(tc.Input)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}

func TestStringToTime(t *testing.T) {
	testCases := map[string]struct {
		Input    string
		Expected time.Time
	}{
		"ok": {
			Input:    "2020-01-01 01:01:01",
			Expected: time.Date(2020, 1, 1, 1, 1, 1, 0, time.Local),
		},
		"invalid format": {
			Input:    "",
			Expected: time.Time{},
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			got := StringToTime(tc.Input)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}

func TestStringToDate(t *testing.T) {
	testCases := map[string]struct {
		Input    string
		Expected time.Time
	}{
		"ok": {
			Input:    "2020-01-01",
			Expected: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
		"invalid format": {
			Input:    "",
			Expected: time.Time{},
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			got := StringToDate(tc.Input)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}

func TestBeginningOfMonth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		expect time.Time
	}{
		{
			name:   "success: 2021-08-01",
			input:  "2021-08-01",
			expect: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:   "success: 2021-08-31",
			input:  "2021-08-31",
			expect: time.Date(2021, time.Month(8), 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:   "invalid format",
			input:  "2021-08-00",
			expect: time.Time{},
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
		input  string
		expect time.Time
	}{
		{
			name:   "success: 2021-08-01",
			input:  "2021-08-01",
			expect: time.Date(2021, time.Month(8), 31, 23, 59, 59, 0, time.Local),
		},
		{
			name:   "success: 2021-08-31",
			input:  "2021-08-31",
			expect: time.Date(2021, time.Month(8), 31, 23, 59, 59, 0, time.Local),
		},
		{
			name:   "invalid format",
			input:  "2021-08-00",
			expect: time.Time{},
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
