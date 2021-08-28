package v1

// 書籍情報
type BookV1Response struct {
	Id           int64                     `json:"id,omitempty"`           // 書籍ID
	Title        string                    `json:"title,omitempty"`        // タイトル
	TitleKana    string                    `json:"titleKana,omitempty"`    // タイトル(かな)
	Description  string                    `json:"description,omitempty"`  // 説明
	Isbn         string                    `json:"isbn,omitempty"`         // ISBN
	Publisher    string                    `json:"publisher,omitempty"`    // 出版社名
	PublishedOn  string                    `json:"published_on,omitempty"` // 出版日
	ThumbnailUrl string                    `json:"thumbnailUrl,omitempty"` // サムネイルURL
	RakutenUrl   string                    `json:"rakutenUrl,omitempty"`   // 楽天ショップURL
	Size         string                    `json:"size,omitempty"`         // 楽天書籍サイズ
	Author       string                    `json:"author,omitempty"`       // 著者名一覧
	AuthorKana   string                    `json:"authorKana,omitempty"`   /// 著者名一覧(かな)
	Bookshelf    *BookV1Response_Bookshelf `json:"bookshelf,omitempty"`    // ユーザーの本棚情報
	CreatedAt    string                    `json:"createdAt,omitempty"`    // 登録日時
	UpdatedAt    string                    `json:"updatedAt,omitempty"`    // 更新日時
}

type BookV1Response_Bookshelf struct {
	Id         int64  `json:"id,omitempty"`         // 本棚ID
	Status     string `json:"status,omitempty"`     // 読書ステータス
	ReadOn     string `json:"readOn,omitempty"`     // 読み終えた日
	Impression string `json:"impression,omitempty"` // 感想
	CreatedAt  string `json:"createdAt,omitempty"`  // 登録日時
	UpdatedAt  string `json:"updatedAt,omitempty"`  // 更新日時
}
