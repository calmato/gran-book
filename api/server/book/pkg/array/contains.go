package array

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrUnsupportedType = errors.New("array: this is unsupported type")

// Contains - 配列に対象の要素が含まれるか
func Contains(items interface{}, target interface{}) (bool, error) {
	switch v := items.(type) {
	case []int:
		for _, item := range v {
			if target == item {
				return true, nil
			}
		}
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
		err := fmt.Errorf("%w, %v", ErrUnsupportedType, reflect.TypeOf(items))
		return false, err
	}

	return false, nil
}
