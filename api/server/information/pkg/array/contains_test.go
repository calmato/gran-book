package array

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := map[string]struct {
		Input struct {
			Items  interface{}
			Target interface{}
		}
		Expected bool
	}{
		"ok_[]int32_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []int32{1, 2, 3},
				Target: int32(2),
			},
			Expected: true,
		},
		"ok_[]int32_not_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []int32{1, 2, 3},
				Target: int32(4),
			},
			Expected: false,
		},
		"ok_[]int64_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []int64{1, 2, 3},
				Target: int64(2),
			},
			Expected: true,
		},
		"ok_[]int64_not_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []int64{1, 2, 3},
				Target: int64(4),
			},
			Expected: false,
		},
		"ok_[]string_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []string{"1", "2", "3"},
				Target: "2",
			},
			Expected: true,
		},
		"ok_[]string_not_contain": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []string{"1", "2", "3"},
				Target: "4",
			},
			Expected: false,
		},
		"ng_other_type": {
			Input: struct {
				Items  interface{}
				Target interface{}
			}{
				Items:  []int{1, 2, 3},
				Target: 2,
			},
			Expected: false,
		},
	}

	for result, tt := range testCases {
		t.Run(result, func(t *testing.T) {
			got, _ := Contains(tt.Input.Items, tt.Input.Target)
			if !reflect.DeepEqual(got, tt.Expected) {
				t.Fatalf("want %#v, but %#v", tt.Expected, got)
			}
		})
	}
}
