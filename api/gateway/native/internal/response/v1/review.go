package v1

// 書籍毎のレビュー情報
type BookReviewResponse struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	User       *BookReviewUser `json:"user"`       // 投稿者
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

type BookReviewUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// ユーザー毎のレビュー情報
type UserReviewResponse struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	Book       *UserReviewBook `json:"book"`       // 書籍情報
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

type UserReviewBook struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// 書籍毎のレビュー一覧
type BookReviewListResponse struct {
	Reviews []*BookReviewListReview `json:"reviewsList"` // レビュー一覧
	Limit   int64                   `json:"limit"`       // 取得上限数
	Offset  int64                   `json:"offset"`      // 取得開始位置
	Total   int64                   `json:"total"`       // 検索一致数
}

type BookReviewListReview struct {
	ID         int64               `json:"id"`         // レビューID
	Impression string              `json:"impression"` // 感想
	User       *BookReviewListUser `json:"user"`       // 投稿者
	CreatedAt  string              `json:"createdAt"`  // 登録日時
	UpdatedAt  string              `json:"updatedAt"`  // 更新日時
}

type BookReviewListUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// ユーザー毎のレビュー一覧
type UserReviewListResponse struct {
	Reviews []*UserReviewListReview `json:"reviewsList"` // レビュー一覧
	Limit   int64                   `json:"limit"`       // 取得上限数
	Offset  int64                   `json:"offset"`      // 取得開始位置
	Total   int64                   `json:"total"`       // 検索一致数
}

type UserReviewListReview struct {
	ID         int64               `json:"id"`         // レビューID
	Impression string              `json:"impression"` // 感想
	Book       *UserReviewListBook `json:"book"`       // 書籍情報
	CreatedAt  string              `json:"createdAt"`  // 登録日時
	UpdatedAt  string              `json:"updatedAt"`  // 更新日時
}

type UserReviewListBook struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
