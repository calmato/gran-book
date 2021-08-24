package array

import (
	"reflect"
	"testing"
)

func TestConvertStrings(t *testing.T) {
	testCases := map[string]struct {
		Input    interface{}
		Expected []string
	}{
		"ok_[]int": {
			Input:    []int{1, 2, 3},
			Expected: []string{"1", "2", "3"},
		},
		"ok_[]int32": {
			Input:    []int32{1, 2, 3},
			Expected: []string{"1", "2", "3"},
		},
		"ok_[]int64": {
			Input:    []int32{1, 2, 3},
			Expected: []string{"1", "2", "3"},
		},
		"ok_[]string": {
			Input:    []string{"1", "2", "3"},
			Expected: []string{"1", "2", "3"},
		},
		"ng_other_type": {
			Input:    []float64{1.0, 2.0, 3.0},
			Expected: []string{},
		},
	}

	for result, tt := range testCases {
		t.Run(result, func(t *testing.T) {
			got, _ := ConvertStrings(tt.Input)
			if !reflect.DeepEqual(got, tt.Expected) {
				t.Fatalf("want %#v, but %#v", tt.Expected, got)
			}
		})
	}
}
