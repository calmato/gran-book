package array

import (
	"fmt"
	"reflect"
	"strconv"
)

// ConvertStrings - 文字列型の配列に変換
func ConvertStrings(items interface{}) ([]string, error) {
	strs := []string{}

	switch v := items.(type) {
	case []int:
		for _, item := range v {
			str := strconv.Itoa(item)
			strs = append(strs, str)
		}
	case []int32:
		for _, item := range v {
			str := strconv.Itoa(int(item))
			strs = append(strs, str)
		}
	case []int64:
		for _, item := range v {
			str := strconv.FormatInt(item, 10)
			strs = append(strs, str)
		}
	case []string:
		strs = v
	default:
		return []string{}, fmt.Errorf("array: %v is an unsupported type", reflect.TypeOf(items))
	}

	return strs, nil
}
