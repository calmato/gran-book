package v2

// 本棚の書籍情報
type BookshelfResponse struct {
	ID           int64               `json:"id"`           // 書籍ID
	Title        string              `json:"title"`        // タイトル
	TitleKana    string              `json:"titleKana"`    // タイトル(かな)
	Description  string              `json:"description"`  // 説明
	Isbn         string              `json:"isbn"`         // ISBN
	Publisher    string              `json:"publisher"`    // 出版社名
	PublishedOn  string              `json:"publishedOn"`  // 出版日
	ThumbnailURL string              `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string              `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string              `json:"size"`         // 楽天書籍サイズ
	Author       string              `json:"author"`       // 著者名一覧
	AuthorKana   string              `json:"author_kana"`  // 著者名一覧(かな)
	Bookshelf    *BookshelfBookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	Reviews      []*BookshelfReview  `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64               `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64               `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64               `json:"reviewTotal"`  // レビュー検索一致件
	CreatedAt    string              `json:"createdAt"`    // 登録日時
	UpdatedAt    string              `json:"updatedAt"`    // 更新日時
}

type BookshelfBookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewID  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

type BookshelfReview struct {
	ID         int64          `json:"id"`         // レビューID
	Impression string         `json:"impression"` // 感想
	User       *BookshelfUser `json:"user"`       // 投稿者
	CreatedAt  string         `json:"createdAt"`  // 登録日時
	UpdatedAt  string         `json:"updatedAt"`  // 更新日時
}

type BookshelfUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// 本棚の書籍一覧
type BookshelfListResponse struct {
	Books  []*BookshelfListBook `json:"booksList"` // 書籍一覧
	Limit  int64                `json:"limit"`     // 取得上限数
	Offset int64                `json:"offset"`    // 取得開始位置
	Total  int64                `json:"total"`     // 検索一致数
}

type BookshelfListBook struct {
	ID           int64                   `json:"id"`           // 書籍ID
	Title        string                  `json:"title"`        // タイトル
	TitleKana    string                  `json:"titleKana"`    // タイトル(かな)
	Description  string                  `json:"description"`  // 説明
	Isbn         string                  `json:"isbn"`         // ISBN
	Publisher    string                  `json:"publisher"`    // 出版社名
	PublishedOn  string                  `json:"publishedOn"`  // 出版日
	ThumbnailURL string                  `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string                  `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                  `json:"size"`         // 楽天書籍サイズ
	Author       string                  `json:"author"`       // 著者名一覧
	AuthorKana   string                  `json:"authorKana"`   // 著者名一覧(かな)
	Bookshelf    *BookshelfListBookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string                  `json:"createdAt"`    // 登録日時
	UpdatedAt    string                  `json:"updatedAt"`    // 更新日時
}

type BookshelfListBookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewID  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}
