package v1

// 書籍情報
type BookResponse struct {
	Id           int64                   `json:"id"`           // 書籍ID
	Title        string                  `json:"title"`        // タイトル
	TitleKana    string                  `json:"titleKana"`    // タイトル(かな)
	Description  string                  `json:"description"`  // 説明
	Isbn         string                  `json:"isbn"`         // ISBN
	Publisher    string                  `json:"publisher"`    // 出版社名
	PublishedOn  string                  `json:"published_on"` // 出版日
	ThumbnailUrl string                  `json:"thumbnailUrl"` // サムネイルURL
	RakutenUrl   string                  `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                  `json:"size"`         // 楽天書籍サイズ
	Author       string                  `json:"author"`       // 著者名一覧
	AuthorKana   string                  `json:"authorKana"`   /// 著者名一覧(かな)
	Bookshelf    *BookResponse_Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string                  `json:"createdAt"`    // 登録日時
	UpdatedAt    string                  `json:"updatedAt"`    // 更新日時
}

type BookResponse_Bookshelf struct {
	Id         int64  `json:"id"`         // 本棚ID
	Status     string `json:"status"`     // 読書ステータス
	ReadOn     string `json:"readOn"`     // 読み終えた日
	Impression string `json:"impression"` // 感想
	CreatedAt  string `json:"createdAt"`  // 登録日時
	UpdatedAt  string `json:"updatedAt"`  // 更新日時
}
