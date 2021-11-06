package response

import "github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"

// 書籍情報
// Deprecated: use v2
type BookResponse struct {
	*entity.Book
}
