package v1

// 書籍毎のレビュー情報
type BookReviewResponse struct {
	ID         int64                    `json:"id"`         // レビューID
	Impression string                   `json:"impression"` // 感想
	User       *BookReviewResponse_User `json:"user"`       // 投稿者
	CreatedAt  string                   `json:"createdAt"`  // 登録日時
	UpdatedAt  string                   `json:"updatedAt"`  // 更新日時
}

type BookReviewResponse_User struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl"` // サムネイルURL
}

// ユーザー毎のレビュー情報
type UserReviewResponse struct {
	ID         int64                    `json:"id"`         // レビューID
	Impression string                   `json:"impression"` // 感想
	Book       *UserReviewResponse_Book `json:"book"`       // 書籍情報
	CreatedAt  string                   `json:"createdAt"`  // 登録日時
	UpdatedAt  string                   `json:"updatedAt"`  // 更新日時
}

type UserReviewResponse_Book struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// 書籍毎のレビュー一覧
type BookReviewListResponse struct {
	Reviews []*BookReviewListResponse_Review `json:"reviewsList"` // レビュー一覧
	Limit   int64                            `json:"limit"`       // 取得上限数
	Offset  int64                            `json:"offset"`      // 取得開始位置
	Total   int64                            `json:"total"`       // 検索一致数
}

type BookReviewListResponse_Review struct {
	ID         int64                        `json:"id"`         // レビューID
	Impression string                       `json:"impression"` // 感想
	User       *BookReviewListResponse_User `json:"user"`       // 投稿者
	CreatedAt  string                       `json:"createdAt"`  // 登録日時
	UpdatedAt  string                       `json:"updatedAt"`  // 更新日時
}

type BookReviewListResponse_User struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// ユーザー毎のレビュー一覧
type UserReviewListResponse struct {
	Reviews []*UserReviewListResponse_Review `json:"reviewsList"` // レビュー一覧
	Limit   int64                            `json:"limit"`       // 取得上限数
	Offset  int64                            `json:"offset"`      // 取得開始位置
	Total   int64                            `json:"total"`       // 検索一致数
}

type UserReviewListResponse_Review struct {
	ID         int64                        `json:"id"`         // レビューID
	Impression string                       `json:"impression"` // 感想
	Book       *UserReviewListResponse_Book `json:"book"`       // 書籍情報
	CreatedAt  string                       `json:"createdAt"`  // 登録日時
	UpdatedAt  string                       `json:"updatedAt"`  // 更新日時
}

type UserReviewListResponse_Book struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
