package entity

import "github.com/calmato/gran-book/api/service/internal/gateway/entity"

type BookReview struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	User       *BookReviewUser `json:"user"`       // 投稿者
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

type BookReviews []*BookReview

type BookReviewUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func NewBookReview(r *entity.Review, u *entity.User) *BookReview {
	return &BookReview{
		ID:         r.Id,
		Impression: r.Impression,
		User:       newBookReviewUser(u),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newBookReviewUser(u *entity.User) *BookReviewUser {
	if u == nil {
		return &BookReviewUser{
			Username: UnknownUserName,
		}
	}

	return &BookReviewUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func NewBookReviews(rs entity.Reviews, um map[string]*entity.User) BookReviews {
	res := make(BookReviews, len(rs))
	for i, r := range rs {
		res[i] = NewBookReview(r, um[r.UserId])
	}
	return res
}

type UserReview struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	Book       *UserReviewBook `json:"book"`       // 書籍情報
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

type UserReviews []*UserReview

type UserReviewBook struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func NewUserReview(r *entity.Review, b *entity.Book) *UserReview {
	return &UserReview{
		ID:         r.Id,
		Impression: r.Impression,
		Book:       newUserReviewBook(b),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newUserReviewBook(b *entity.Book) *UserReviewBook {
	return &UserReviewBook{
		ID:           b.Id,
		Title:        b.Title,
		ThumbnailURL: b.ThumbnailUrl,
	}
}

func NewUserReviews(rs entity.Reviews, bm map[int64]*entity.Book) UserReviews {
	res := make(UserReviews, len(rs))
	for i, r := range rs {
		res[i] = NewUserReview(r, bm[r.BookId])
	}
	return res
}
