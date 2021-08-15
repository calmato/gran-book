package v2

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
	Reviews      []*BookshelfReview  `json:"reviews"`
	ReviewLimit  int64               `json:"reviewLimit"`
	ReviewOffset int64               `json:"reviewLOffset"`
	ReviewTotal  int64               `json:"reviewTotal"`
}

type BookshelfBookshelf struct {
	ReviewID  int64  `json:"reviewId,omitempty"`
	Status    string `json:"status"`
	ReadOn    string `json:"readOn"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type BookshelfReview struct {
	ID         int64          `json:"id"`
	Impression string         `json:"impression"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	User       *BookshelfUser `json:"user"`
}

type BookshelfUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
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
	ReviewID  int64  `json:"reviewId,omitempty"`
	Status    string `json:"status"`
	ReadOn    string `json:"readOn"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
