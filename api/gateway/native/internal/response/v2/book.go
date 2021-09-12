package v2

// 書籍情報
type BookResponse struct {
	ID           int64         `json:"id"`           // 書籍ID
	Title        string        `json:"title"`        // タイトル
	TitleKana    string        `json:"titleKana"`    // タイトル(かな)
	Description  string        `json:"description"`  // 説明
	Isbn         string        `json:"isbn"`         // ISBN
	Publisher    string        `json:"publisher"`    // 出版社名
	PublishedOn  string        `json:"publishedOn"`  // 出版日
	ThumbnailURL string        `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string        `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string        `json:"size"`         // 楽天書籍サイズ
	Author       string        `json:"author"`       // 著者名一覧
	AuthorKana   string        `json:"authorKana"`   // 著者名一覧(かな)
	Reviews      []*BookReview `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64         `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64         `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64         `json:"reviewTotal"`  // レビュー検索一致件数
	CreatedAt    string        `json:"createdAt"`    // 登録日時
	UpdatedAt    string        `json:"updatedAt"`    // 更新日時
}

type BookReview struct {
	ID         int64     `json:"id"`         // レビューID
	Impression string    `json:"impression"` // 感想
	User       *BookUser `json:"user"`       // 投稿者
	CreatedAt  string    `json:"createdAt"`  // 登録日時
	UpdatedAt  string    `json:"updatedAt"`  // 更新日時
}

type BookUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
