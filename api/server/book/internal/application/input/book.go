package input

// CreateBook - 書籍登録のリクエスト
type CreateBook struct {
	Title        string                `json:"title" validate:"required,max=32"`
	Description  string                `json:"description" validate:"max=1000"`
	Isbn         string                `json:"isbn" validate:"required"`
	ThumbnailURL string                `json:"thumbnailUrl"`
	Version      string                `json:"version"`
	Publisher    string                `json:"publisher" validate:"max=32"`
	PublishedOn  string                `json:"publishedOn"`
	Authors      []*CreateBookAuthor   `json:"authors"`
	Categories   []*CreateBookCategory `json:"categories"`
}

type CreateBookAuthor struct {
	Name string `json:"name" validate:"required,max=32"`
}

type CreateBookCategory struct {
	Name string `json:"name" validate:"required,max=32"`
}
