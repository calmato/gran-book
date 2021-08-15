package v1

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

type UserReviewResponse struct {
	ID         int64           `json:"id"`
	Impression string          `json:"impression"`
	CreatedAt  string          `json:"createdAt"`
	UpdatedAt  string          `json:"updatedAt"`
	Book       *UserReviewBook `json:"book"`
}

type UserReviewBook struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type UserReviewListResponse struct {
	Reviews []*UserReviewListReview `json:"reviews"`
	Limit   int64                   `json:"limit"`
	Offset  int64                   `json:"offset"`
	Total   int64                   `json:"total"`
}

type UserReviewListReview struct {
	ID         int64               `json:"id"`
	Impression string              `json:"impression"`
	CreatedAt  string              `json:"createdAt"`
	UpdatedAt  string              `json:"updatedAt"`
	Book       *UserReviewListBook `json:"book"`
}

type UserReviewListBook struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
