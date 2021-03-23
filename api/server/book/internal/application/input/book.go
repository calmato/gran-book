package input

// BookItem - 書籍登録/更新のリクエスト
type BookItem struct {
	Title        string          `json:"title" validate:"required,max=64"`
	Description  string          `json:"description" validate:"max=1600"`
	Isbn         string          `json:"isbn" validate:"required"`
	ThumbnailURL string          `json:"thumbnailUrl"`
	Version      string          `json:"version"`
	Publisher    string          `json:"publisher" validate:"max=32"`
	PublishedOn  string          `json:"publishedOn"`
	Authors      []*BookAuthor   `json:"authors" validate:"omitempty,dive"`
	Categories   []*BookCategory `json:"categories" validate:"omitempty,dive"`
}

// CreateAndUpdateBooks - 書籍一括登録/更新のリクエスト
type CreateAndUpdateBooks struct {
	Books []*BookItem `json:"books" validate:"dive"`
}

type BookAuthor struct {
	Name string `json:"name" validate:"required,max=32"`
}

type BookCategory struct {
	Name string `json:"name" validate:"required,max=32"`
}
