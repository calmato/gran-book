package response

import "github.com/calmato/gran-book/api/internal/gateway/native/v2/entity"

// 書籍情報
type BookResponse struct {
	*entity.Book
}
