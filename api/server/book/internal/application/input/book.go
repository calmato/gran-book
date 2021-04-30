package input

// ListOnBookshelf - 本棚ないの書籍一覧取得のリクエスト
type ListBookshelf struct {
	UserID string `json:"userId" validate:"required"`
	Limit  int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset int    `json:"offset" validate:"gte=0"`
}

// Book - 書籍登録/更新のリクエスト
type Book struct {
	Title          string        `json:"title" validate:"required,max=64"`
	TitleKana      string        `json:"titleKana" validate:"required,max=128"`
	Description    string        `json:"description" validate:"omitempty,max=2000"`
	Isbn           string        `json:"isbn" validate:"required,max=16"`
	Publisher      string        `json:"publisher" validate:"required,max=32"`
	PublishedOn    string        `json:"publishedOn" validate:"required"`
	ThumbnailURL   string        `json:"thumbnailUrl" validate:"omitempty"`
	RakutenURL     string        `json:"rakutenUrl" validate:"omitempty"`
	RakutenGenreID string        `json:"rakutenGenreId" validate:"omitempty"`
	Authors        []*BookAuthor `json:"authors" validate:"omitempty"`
}

// BookAuthor - 著者のリクエスト
type BookAuthor struct {
	Name     string `json:"name" validate:"required,max=32"`
	NameKana string `json:"nameKana" validate:"required,max=64"`
}

// Bookshelf - 本棚への書籍登録/更新のリクエスト
type Bookshelf struct {
	UserID string `json:"userId" validate:"required"`
	BookID int    `json:"bookId" validate:"required,gte=1"`
	Status int    `json:"status" validate:"required,gte=0,lte=5"`
	ReadOn string `json:"readOn" validate:"omitempty"`
}
