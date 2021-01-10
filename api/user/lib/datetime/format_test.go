package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	testCases := map[string]struct {
		Input    time.Time
		Expected string
	}{
		"ok": {
			Input:    time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC),
			Expected: "2020-01-01 01:01:01",
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

func TestStringToTime(t *testing.T) {
	testCases := map[string]struct {
		Input    string
		Expected time.Time
	}{
		"ok": {
			Input:    "2020-01-01 01:01:01",
			Expected: time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC),
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
