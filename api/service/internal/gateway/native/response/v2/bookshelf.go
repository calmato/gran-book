package v2

import (
	"github.com/calmato/gran-book/api/service/internal/gateway/native/entity"
)

// 本棚の書籍情報
type BookshelfResponse struct {
	*entity.BookOnBookshelf
}

// 本棚の書籍一覧
type BookshelfListResponse struct {
	Books  entity.BooksOnBookshelf `json:"booksList"` // 書籍一覧
	Limit  int64                   `json:"limit"`     // 取得上限数
	Offset int64                   `json:"offset"`    // 取得開始位置
	Total  int64                   `json:"total"`     // 検索一致数
}
