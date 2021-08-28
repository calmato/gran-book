package v1

// 書籍登録
type CreateBookRequest struct {
	Title          string `json:"title"`          // タイトル
	TitleKana      string `json:"titleKana"`      // タイトル(かな)
	ItemCaption    string `json:"itemCaption"`    // 説明
	Isbn           string `json:"isbn"`           // ISBN
	PublisherName  string `json:"publisherName"`  // 出版社名
	SalesDate      string `json:"salesDate"`      // 出版日
	SmallImageURL  string `json:"smallImageUrl"`  // サムネイルURL(Sサイズ)
	MediumImageURL string `json:"mediumImageUrl"` // サムネイルURL(Mサイズ)
	LargeImageURL  string `json:"largeImageUrl"`  // サムネイルURL(Lサイズ)
	ItemURL        string `json:"itemUrl"`        // 楽天ショップURL
	Size           string `json:"size"`           // 楽天書籍サイズ
	BooksGenreID   string `json:"booksGenreId"`   // 楽天書籍ジャンルID
	Author         string `json:"author"`         // 著者名一覧
	AuthorKana     string `json:"authorKana"`     // 著者名一覧(かな)
}

// 書籍更新
type UpdateBookRequest struct {
	Title          string `json:"title"`          // タイトル
	TitleKana      string `json:"titleKana"`      // タイトル(かな)
	ItemCaption    string `json:"itemCaption"`    // 説明
	Isbn           string `json:"isbn"`           // ISBN
	PublisherName  string `json:"publisherName"`  // 出版社名
	SalesDate      string `json:"salesDate"`      // 出版日
	SmallImageURL  string `json:"smallImageUrl"`  // サムネイルURL(Sサイズ)
	MediumImageURL string `json:"mediumImageUrl"` // サムネイルURL(Mサイズ)
	LargeImageURL  string `json:"largeImageUrl"`  // サムネイルURL(Lサイズ)
	ItemURL        string `json:"itemUrl"`        // 楽天ショップURL
	Size           string `json:"size"`           // 楽天書籍サイズ
	BooksGenreID   string `json:"booksGenreId"`   // 楽天書籍ジャンルID
	Author         string `json:"author"`         // 著者名一覧
	AuthorKana     string `json:"authorKana"`     // 著者名一覧(かな)
}
