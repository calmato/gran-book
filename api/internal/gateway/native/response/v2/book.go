package v2

import "github.com/calmato/gran-book/api/internal/gateway/native/entity"

// 書籍情報
type BookResponse struct {
	*entity.Book
}
