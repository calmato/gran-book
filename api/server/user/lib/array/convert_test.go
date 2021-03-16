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
			Input:    []int{1, 2, 3},
			Expected: []string{},
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			got, _ := ConvertStrings(tc.Input)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}
