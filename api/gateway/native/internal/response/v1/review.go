package v1

// 書籍毎のレビュー情報
type BookReviewV1Response struct {
	Id         int64                      `json:"id,omitempty"`         // レビューID
	Impression string                     `json:"impression,omitempty"` // 感想
	User       *BookReviewV1Response_User `json:"user,omitempty"`       // 投稿者
	CreatedAt  string                     `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string                     `json:"updatedAt,omitempty"`  // 更新日時
}

type BookReviewV1Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// ユーザー毎のレビュー情報
type UserReviewV1Response struct {
	Id         int64                      `json:"id,omitempty"`         // レビューID
	Impression string                     `json:"impression,omitempty"` // 感想
	Book       *UserReviewV1Response_Book `json:"book,omitempty"`       // 書籍情報
	CreatedAt  string                     `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string                     `json:"updatedAt,omitempty"`  // 更新日時
}

type UserReviewV1Response_Book struct {
	Id           int64  `json:"id,omitempty"`           // 書籍ID
	Title        string `json:"title,omitempty"`        // タイトル
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// 書籍毎のレビュー一覧
type BookReviewListV1Response struct {
	Reviews []*BookReviewListV1Response_Review `json:"reviewsList,omitempty"` // レビュー一覧
	Limit   int64                              `json:"limit,omitempty"`       // 取得上限数
	Offset  int64                              `json:"offset,omitempty"`      // 取得開始位置
	Total   int64                              `json:"total,omitempty"`       // 検索一致数
}

type BookReviewListV1Response_Review struct {
	Id         int64                          `json:"id,omitempty"`         // レビューID
	Impression string                         `json:"impression,omitempty"` // 感想
	User       *BookReviewListV1Response_User `json:"user,omitempty"`       // 投稿者
	CreatedAt  string                         `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string                         `json:"updatedAt,omitempty"`  // 更新日時
}

type BookReviewListV1Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// ユーザー毎のレビュー一覧
type UserReviewListV1Response struct {
	Reviews []*UserReviewListV1Response_Review `json:"reviewsList,omitempty"` // レビュー一覧
	Limit   int64                              `json:"limit,omitempty"`       // 取得上限数
	Offset  int64                              `json:"offset,omitempty"`      // 取得開始位置
	Total   int64                              `json:"total,omitempty"`       // 検索一致数
}

type UserReviewListV1Response_Review struct {
	Id         int64                          `json:"id,omitempty"`         // レビューID
	Impression string                         `json:"impression,omitempty"` // 感想
	Book       *UserReviewListV1Response_Book `json:"book,omitempty"`       // 書籍情報
	CreatedAt  string                         `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string                         `json:"updatedAt,omitempty"`  // 更新日時
}

type UserReviewListV1Response_Book struct {
	Id           int64  `json:"id,omitempty"`           // 書籍ID
	Title        string `json:"title,omitempty"`        // タイトル
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}
