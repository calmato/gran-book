package input

// Book - 書籍登録/更新のリクエスト
type Book struct {
	Title        string   `json:"title" validate:"required,max=64"`
	Description  string   `json:"description" validate:"omitempty,max=1600"`
	Isbn         string   `json:"isbn" validate:"required,max=32"`
	ThumbnailURL string   `json:"thumbnailUrl" validate:"omitempty"`
	Version      string   `json:"version" validate:"required"`
	Publisher    string   `json:"publisher" validate:"omitempty,max=32"`
	PublishedOn  string   `json:"publishedOn" validate:"omitempty"`
	Authors      []string `json:"authors" validate:"dive,required,max=32"`
	Categories   []string `json:"categories" validate:"dive,required,max=32"`
}

// Bookshelf - 本棚への書籍登録/更新のリクエスト
type Bookshelf struct {
	UserID     string `json:"userId" validate:"required"`
	BookID     int    `json:"bookId" validate:"required,gte=1"`
	Status     int    `json:"status" validate:"required,gte=0,lte=5"`
	Impression string `json:"impression" validate:"omitempty,max=1000"`
	ReadOn     string `json:"readOn" validate:"omitempty"`
}

// CreateAndUpdateBooks - 書籍一括登録/更新のリクエスト
type CreateAndUpdateBooks struct {
	Books []*Book `json:"books" validate:"omitempty,dive"`
}
