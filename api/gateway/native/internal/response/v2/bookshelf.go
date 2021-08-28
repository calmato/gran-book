package v2

// 本棚の書籍情報
type BookshelfV2Response struct {
	Id           int64                          `json:"id,omitempty"`           // 書籍ID
	Title        string                         `json:"title,omitempty"`        // タイトル
	TitleKana    string                         `json:"titleKana,omitempty"`    // タイトル(かな)
	Description  string                         `json:"description,omitempty"`  // 説明
	Isbn         string                         `json:"isbn,omitempty"`         // ISBN
	Publisher    string                         `json:"publisher,omitempty"`    // 出版社名
	PublishedOn  string                         `json:"publishedOn,omitempty"`  // 出版日
	ThumbnailUrl string                         `json:"thumbnailUrl,omitempty"` // サムネイルURL
	RakutenUrl   string                         `json:"rakutenUrl,omitempty"`   // 楽天ショップURL
	Size         string                         `json:"size,omitempty"`         // 楽天書籍サイズ
	Author       string                         `json:"author,omitempty"`       // 著者名一覧
	AuthorKana   string                         `json:"author_kana,omitempty"`  /// 著者名一覧(かな)
	Bookshelf    *BookshelfV2Response_Bookshelf `json:"bookshelf,omitempty"`    // ユーザーの本棚情報
	Reviews      []*BookshelfV2Response_Review  `json:"reviewsList,omitempty"`  // レビュー一覧
	ReviewLimit  int64                          `json:"reviewLimit,omitempty"`  // レビュー取得上限
	ReviewOffset int64                          `json:"reviewOffset,omitempty"` // レビュー取得開始位置
	ReviewTotal  int64                          `json:"reviewTotal,omitempty"`  // レビュー検索一致件
	CreatedAt    string                         `json:"createdAt,omitempty"`    // 登録日時
	UpdatedAt    string                         `json:"updatedAt,omitempty"`    // 更新日時
}

type BookshelfV2Response_Bookshelf struct {
	Status    string `json:"status,omitempty"`    // 読書ステータス
	ReadOn    string `json:"readOn,omitempty"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId,omitempty"`  // レビューID
	CreatedAt string `json:"createdAt,omitempty"` // 登録日時
	UpdatedAt string `json:"updatedAt,omitempty"` // 更新日時
}

type BookshelfV2Response_Review struct {
	Id         int64                     `json:"id,omitempty"`         // レビューID
	Impression string                    `json:"impression,omitempty"` // 感想
	User       *BookshelfV2Response_User `json:"user,omitempty"`       // 投稿者
	CreatedAt  string                    `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string                    `json:"updatedAt,omitempty"`  // 更新日時
}

type BookshelfV2Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// 本棚の書籍一覧
type BookshelfListV2Response struct {
	Books  []*BookshelfListV2Response_Book `json:"booksList,omitempty"` // 書籍一覧
	Limit  int64                           `json:"limit,omitempty"`     // 取得上限数
	Offset int64                           `json:"offset,omitempty"`    // 取得開始位置
	Total  int64                           `json:"total,omitempty"`     // 検索一致数
}

type BookshelfListV2Response_Book struct {
	Id           int64                              `json:"id,omitempty"`           // 書籍ID
	Title        string                             `json:"title,omitempty"`        // タイトル
	TitleKana    string                             `json:"titleKana,omitempty"`    // タイトル(かな)
	Description  string                             `json:"description,omitempty"`  // 説明
	Isbn         string                             `json:"isbn,omitempty"`         // ISBN
	Publisher    string                             `json:"publisher,omitempty"`    // 出版社名
	PublishedOn  string                             `json:"publishedOn,omitempty"`  // 出版日
	ThumbnailUrl string                             `json:"thumbnailUrl,omitempty"` // サムネイルURL
	RakutenUrl   string                             `json:"rakutenUrl,omitempty"`   // 楽天ショップURL
	Size         string                             `json:"size,omitempty"`         // 楽天書籍サイズ
	Author       string                             `json:"author,omitempty"`       // 著者名一覧
	AuthorKana   string                             `json:"authorKana,omitempty"`   /// 著者名一覧(かな)
	Bookshelf    *BookshelfListV2Response_Bookshelf `json:"bookshelf,omitempty"`    // ユーザーの本棚情報
	CreatedAt    string                             `json:"createdAt,omitempty"`    // 登録日時
	UpdatedAt    string                             `json:"updatedAt,omitempty"`    // 更新日時
}

type BookshelfListV2Response_Bookshelf struct {
	Status    string `json:"status,omitempty"`    // 読書ステータス
	ReadOn    string `json:"readOn,omitempty"`    // 読み終えた日
	ReviewId  int64  `json:"reviewId,omitempty"`  // レビューID
	CreatedAt string `json:"createdAt,omitempty"` // 登録日時
	UpdatedAt string `json:"updatedAt,omitempty"` // 更新日時
}
