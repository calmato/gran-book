package v1

type BookshelfResponse struct {
	ID           int64               `json:"id"`
	Title        string              `json:"title"`
	TitleKana    string              `json:"titleKana"`
	Description  string              `json:"description"`
	Isbn         string              `json:"isbn"`
	Publisher    string              `json:"publisher"`
	PublishedOn  string              `json:"publishedOn"`
	ThumbnailURL string              `json:"thumbnailUrl"`
	RakutenURL   string              `json:"rakutenUrl"`
	Size         string              `json:"size"`
	Author       string              `json:"author"`
	AuthorKana   string              `json:"authorKana"`
	CreatedAt    string              `json:"createdAt"`
	UpdatedAt    string              `json:"updatedAt"`
	Bookshelf    *BookshelfBookshelf `json:"bookshelf,omitempty"`
}

type BookshelfBookshelf struct {
	ID         int64  `json:"id"`
	Status     string `json:"status"`
	ReadOn     string `json:"readOn"`
	Impression string `json:"impression"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type BookshelfListResponse struct {
	Books  []*BookshelfListBook `json:"books"`
	Limit  int64                `json:"limit"`
	Offset int64                `json:"offset"`
	Total  int64                `json:"total"`
}

type BookshelfListBook struct {
	ID           int64                   `json:"id"`
	Title        string                  `json:"title"`
	TitleKana    string                  `json:"titleKana"`
	Description  string                  `json:"description"`
	Isbn         string                  `json:"isbn"`
	Publisher    string                  `json:"publisher"`
	PublishedOn  string                  `json:"publishedOn"`
	ThumbnailURL string                  `json:"thumbnailUrl"`
	RakutenURL   string                  `json:"rakutenUrl"`
	Size         string                  `json:"size"`
	Author       string                  `json:"author"`
	AuthorKana   string                  `json:"authorKana"`
	CreatedAt    string                  `json:"createdAt"`
	UpdatedAt    string                  `json:"updatedAt"`
	Bookshelf    *BookshelfListBookshelf `json:"bookshelf,omitempty"`
}

type BookshelfListBookshelf struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	ReadOn    string `json:"readOn"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
