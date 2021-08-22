package array

import (
	"reflect"

	"golang.org/x/xerrors"
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
		err := xerrors.Errorf("%v is an unsupported type.", reflect.TypeOf(items))
		return false, err
	}

	return false, nil
}
