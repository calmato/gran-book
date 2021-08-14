package v1

type BookResponse struct {
	ID           int64          `json:"id"`
	Title        string         `json:"title"`
	TitleKana    string         `json:"titleKana"`
	Description  string         `json:"description"`
	Isbn         string         `json:"isbn"`
	Publisher    string         `json:"publisher"`
	PublishedOn  string         `json:"publishedOn"`
	ThumbnailURL string         `json:"thumbnailUrl"`
	RakutenURL   string         `json:"rakutenUrl"`
	Size         string         `json:"size"`
	Author       string         `json:"author"`
	AuthorKana   string         `json:"authorKana"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
	Bookshelf    *BookBookshelf `json:"bookshelf,omitempty"`
}

type BookBookshelf struct {
	ID         int64  `json:"id"`
	Status     string `json:"status"`
	ReadOn     string `json:"readOn"`
	Impression string `json:"impression"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type BookReviewResponse struct {
	ID         int64           `json:"id"`
	Impression string          `json:"impression"`
	CreatedAt  string          `json:"createdAt"`
	UpdatedAt  string          `json:"updatedAt"`
	User       *BookReviewUser `json:"user,omitempty"`
}

type BookReviewUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type BookReviewListResponse struct {
	Reviews []*BookReviewListReview `json:"reviews"`
	Limit   int64                   `json:"limit"`
	Offset  int64                   `json:"offset"`
	Total   int64                   `json:"total"`
}

type BookReviewListReview struct {
	ID         int64               `json:"id"`
	Impression string              `json:"impression"`
	CreatedAt  string              `json:"createdAt"`
	UpdatedAt  string              `json:"updatedAt"`
	User       *BookReviewListUser `json:"user,omitempty"`
}

type BookReviewListUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
