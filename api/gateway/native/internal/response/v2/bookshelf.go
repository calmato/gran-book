package v2

// 本棚の書籍情報
type BookshelfResponse struct {
	Id           int64                        `json:"id"`           // 書籍ID
	Title        string                       `json:"title"`        // タイトル
	TitleKana    string                       `json:"titleKana"`    // タイトル(かな)
	Description  string                       `json:"description"`  // 説明
	Isbn         string                       `json:"isbn"`         // ISBN
	Publisher    string                       `json:"publisher"`    // 出版社名
	PublishedOn  string                       `json:"publishedOn"`  // 出版日
	ThumbnailUrl string                       `json:"thumbnailUrl"` // サムネイルURL
	RakutenUrl   string                       `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                       `json:"size"`         // 楽天書籍サイズ
	Author       string                       `json:"author"`       // 著者名一覧
	AuthorKana   string                       `json:"author_kana"`  /// 著者名一覧(かな)
	Bookshelf    *BookshelfResponse_Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	Reviews      []*BookshelfResponse_Review  `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64                        `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64                        `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64                        `json:"reviewTotal"`  // レビュー検索一致件
	CreatedAt    string                       `json:"createdAt"`    // 登録日時
	UpdatedAt    string                       `json:"updatedAt"`    // 更新日時
}

type BookshelfResponse_Bookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

type BookshelfResponse_Review struct {
	Id         int64                   `json:"id"`         // レビューID
	Impression string                  `json:"impression"` // 感想
	User       *BookshelfResponse_User `json:"user"`       // 投稿者
	CreatedAt  string                  `json:"createdAt"`  // 登録日時
	UpdatedAt  string                  `json:"updatedAt"`  // 更新日時
}

type BookshelfResponse_User struct {
	Id           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl"` // サムネイルURL
}

// 本棚の書籍一覧
type BookshelfListResponse struct {
	Books  []*BookshelfListResponse_Book `json:"booksList"` // 書籍一覧
	Limit  int64                         `json:"limit"`     // 取得上限数
	Offset int64                         `json:"offset"`    // 取得開始位置
	Total  int64                         `json:"total"`     // 検索一致数
}

type BookshelfListResponse_Book struct {
	Id           int64                            `json:"id"`           // 書籍ID
	Title        string                           `json:"title"`        // タイトル
	TitleKana    string                           `json:"titleKana"`    // タイトル(かな)
	Description  string                           `json:"description"`  // 説明
	Isbn         string                           `json:"isbn"`         // ISBN
	Publisher    string                           `json:"publisher"`    // 出版社名
	PublishedOn  string                           `json:"publishedOn"`  // 出版日
	ThumbnailUrl string                           `json:"thumbnailUrl"` // サムネイルURL
	RakutenUrl   string                           `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                           `json:"size"`         // 楽天書籍サイズ
	Author       string                           `json:"author"`       // 著者名一覧
	AuthorKana   string                           `json:"authorKana"`   /// 著者名一覧(かな)
	Bookshelf    *BookshelfListResponse_Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string                           `json:"createdAt"`    // 登録日時
	UpdatedAt    string                           `json:"updatedAt"`    // 更新日時
}

type BookshelfListResponse_Bookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}
