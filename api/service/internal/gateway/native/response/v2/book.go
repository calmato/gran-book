package v2

import "github.com/calmato/gran-book/api/service/internal/gateway/native/entity"

// 書籍情報
type BookResponse struct {
	*entity.Book
}
