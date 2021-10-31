package array

import (
	"fmt"
	"reflect"
)

// Contains - 配列に対象の要素が含まれるか
func Contains(items interface{}, target interface{}) (bool, error) {
	switch v := items.(type) {
	case []int32:
		for _, item := range v {
			if target == item {
				return true, nil
			}
		}
	case []int64:
		for _, item := range v {
			if target == item {
				return true, nil
			}
		}
	case []string:
		for _, item := range v {
			if target == item {
				return true, nil
			}
		}
	default:
		return false, fmt.Errorf("array: %v is an unsupported type", reflect.TypeOf(items))
	}

	return false, nil
}
