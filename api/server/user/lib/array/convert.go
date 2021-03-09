package array

import (
	"reflect"
	"strconv"

	"golang.org/x/xerrors"
)

// ConvertStrings - 文字列型の配列に変換
func ConvertStrings(items interface{}) ([]string, error) {
	strs := []string{}
	switch v := items.(type) {
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
		err := xerrors.Errorf("%v is an unsupported type.", reflect.TypeOf(items))
		return []string{}, err
	}

	return strs, nil
}
