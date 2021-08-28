package v1

// 書籍登録
type CreateBookV1Request struct {
	Title          string `json:"title,omitempty"`          // タイトル
	TitleKana      string `json:"titleKana,omitempty"`      // タイトル(かな)
	ItemCaption    string `json:"itemCaption,omitempty"`    // 説明
	Isbn           string `json:"isbn,omitempty"`           // ISBN
	PublisherName  string `json:"publisherName,omitempty"`  // 出版社名
	SalesDate      string `json:"salesDate,omitempty"`      // 出版日
	SmallImageUrl  string `json:"smallImageUrl,omitempty"`  // サムネイルURL(Sサイズ)
	MediumImageUrl string `json:"mediumImageUrl,omitempty"` // サムネイルURL(Mサイズ)
	LargeImageUrl  string `json:"largeImageUrl,omitempty"`  // サムネイルURL(Lサイズ)
	ItemUrl        string `json:"itemUrl,omitempty"`        // 楽天ショップURL
	Size           string `json:"size,omitempty"`           // 楽天書籍サイズ
	BooksGenreId   string `json:"booksGenreId,omitempty"`   // 楽天書籍ジャンルID
	Author         string `json:"author,omitempty"`         // 著者名一覧
	AuthorKana     string `json:"authorKana,omitempty"`     // 著者名一覧(かな)
}

// 書籍更新
type UpdateBookV1Request struct {
	Title          string `json:"title,omitempty"`          // タイトル
	TitleKana      string `json:"titleKana,omitempty"`      // タイトル(かな)
	ItemCaption    string `json:"itemCaption,omitempty"`    // 説明
	Isbn           string `json:"isbn,omitempty"`           // ISBN
	PublisherName  string `json:"publisherName,omitempty"`  // 出版社名
	SalesDate      string `json:"salesDate,omitempty"`      // 出版日
	SmallImageUrl  string `json:"smallImageUrl,omitempty"`  // サムネイルURL(Sサイズ)
	MediumImageUrl string `json:"mediumImageUrl,omitempty"` // サムネイルURL(Mサイズ)
	LargeImageUrl  string `json:"largeImageUrl,omitempty"`  // サムネイルURL(Lサイズ)
	ItemUrl        string `json:"itemUrl,omitempty"`        // 楽天ショップURL
	Size           string `json:"size,omitempty"`           // 楽天書籍サイズ
	BooksGenreId   string `json:"booksGenreId,omitempty"`   // 楽天書籍ジャンルID
	Author         string `json:"author,omitempty"`         // 著者名一覧
	AuthorKana     string `json:"authorKana,omitempty"`     // 著者名一覧(かな)
}
