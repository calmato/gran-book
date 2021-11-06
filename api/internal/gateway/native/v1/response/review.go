package response

import "github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"

// 書籍毎のレビュー情報
type BookReviewResponse struct {
	*entity.BookReview
}

// 書籍毎のレビュー一覧
type BookReviewListResponse struct {
	Reviews entity.BookReviews `json:"reviewsList"` // レビュー一覧
	Limit   int64              `json:"limit"`       // 取得上限数
	Offset  int64              `json:"offset"`      // 取得開始位置
	Total   int64              `json:"total"`       // 検索一致数
}

// ユーザー毎のレビュー情報
type UserReviewResponse struct {
	*entity.UserReview
}

// ユーザー毎のレビュー一覧
type UserReviewListResponse struct {
	Reviews entity.UserReviews `json:"reviewsList"` // レビュー一覧
	Limit   int64              `json:"limit"`       // 取得上限数
	Offset  int64              `json:"offset"`      // 取得開始位置
	Total   int64              `json:"total"`       // 検索一致数
}
