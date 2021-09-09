package conv

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrStringConvert = errors.New("failed to convert from CamelCase to snake_case")

// CamelToSnake - キャメルケースからスネークケースへの変換
func CamelToSnake(camel string) (string, error) {
	b := &strings.Builder{}
	b.Grow(len(camel))

	for i, r := range camel {
		// 初めの文字 -> 小文字に変換
		if i == 0 {
			if _, err := b.WriteRune(unicode.ToLower(r)); err != nil {
				return "", fmt.Errorf("conv: %w", ErrStringConvert)
			}
			continue
		}

		// 大文字 -> _ + 小文字 に変換
		if unicode.IsUpper(r) {
			if _, err := b.WriteRune('_'); err != nil {
				return "", fmt.Errorf("conv: %w", ErrStringConvert)
			}
			if _, err := b.WriteRune(unicode.ToLower(r)); err != nil {
				return "", fmt.Errorf("conv: %w", ErrStringConvert)
			}
			continue
		}

		// 小文字 -> そのまま
		if _, err := b.WriteRune(r); err != nil {
			return "", fmt.Errorf("conv: %w", ErrStringConvert)
		}
	}

	return b.String(), nil
}
