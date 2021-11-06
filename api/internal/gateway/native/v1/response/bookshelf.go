package response

import "github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"

// 本棚の書籍情報
// Deprecated: use v2
type BookshelfResponse struct {
	*entity.Book
}

// 本棚の書籍一覧
// Deprecated: use v2
type BookshelfListResponse struct {
	Books  entity.Books `json:"booksList"` // 書籍一覧
	Limit  int64        `json:"limit"`     // 取得上限数
	Offset int64        `json:"offset"`    // 取得開始位置
	Total  int64        `json:"total"`     // 検索一致数
}
