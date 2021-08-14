package v2

type BookResponse struct {
	ID           int64         `json:"id"`
	Title        string        `json:"title"`
	TitleKana    string        `json:"titleKana"`
	Description  string        `json:"description"`
	Isbn         string        `json:"isbn"`
	Publisher    string        `json:"publisher"`
	PublishedOn  string        `json:"publishedOn"`
	ThumbnailURL string        `json:"thumbnailUrl"`
	RakutenURL   string        `json:"rakutenUrl"`
	Size         string        `json:"size"`
	Author       string        `json:"author"`
	AuthorKana   string        `json:"authorKana"`
	CreatedAt    string        `json:"createdAt"`
	UpdatedAt    string        `json:"updatedAt"`
	Reviews      []*BookReview `json:"reviews"`
	ReviewLimit  int64         `json:"reviewLimit"`
	ReviewOffset int64         `json:"reviewLOffset"`
	ReviewTotal  int64         `json:"reviewTotal"`
}

type BookReview struct {
	ID         int64     `json:"id"`
	Impression string    `json:"impression"`
	CreatedAt  string    `json:"createdAt"`
	UpdatedAt  string    `json:"updatedAt"`
	User       *BookUser `json:"user"`
}

type BookUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
