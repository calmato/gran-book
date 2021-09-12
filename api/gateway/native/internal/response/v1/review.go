package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

// 書籍毎のレビュー情報
type BookReviewResponse struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	User       *bookReviewUser `json:"user"`       // 投稿者
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

func NewBookReviewResponse(r *entity.Review, u *entity.User) *BookReviewResponse {
	return &BookReviewResponse{
		ID:         r.Id,
		Impression: r.Impression,
		User:       newBookReviewUser(u),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

type bookReviewUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newBookReviewUser(u *entity.User) *bookReviewUser {
	if u == nil {
		return &bookReviewUser{
			Username: "unknown",
		}
	}

	return &bookReviewUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

// ユーザー毎のレビュー情報
type UserReviewResponse struct {
	ID         int64           `json:"id"`         // レビューID
	Impression string          `json:"impression"` // 感想
	Book       *userReviewBook `json:"book"`       // 書籍情報
	CreatedAt  string          `json:"createdAt"`  // 登録日時
	UpdatedAt  string          `json:"updatedAt"`  // 更新日時
}

func NewUserReviewResponse(r *entity.Review, b *entity.Book) *UserReviewResponse {
	return &UserReviewResponse{
		ID:         r.Id,
		Impression: r.Impression,
		Book:       newUserReviewBook(b),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

type userReviewBook struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newUserReviewBook(b *entity.Book) *userReviewBook {
	return &userReviewBook{
		ID:           b.Id,
		Title:        b.Title,
		ThumbnailURL: b.ThumbnailUrl,
	}
}

// 書籍毎のレビュー一覧
type BookReviewListResponse struct {
	Reviews []*bookReviewListReview `json:"reviewsList"` // レビュー一覧
	Limit   int64                   `json:"limit"`       // 取得上限数
	Offset  int64                   `json:"offset"`      // 取得開始位置
	Total   int64                   `json:"total"`       // 検索一致数
}

func NewBookReviewListResponse(
	rs entity.Reviews, um map[string]*entity.User, limit, offset, total int64,
) *BookReviewListResponse {
	return &BookReviewListResponse{
		Reviews: newBookReviewListReviews(rs, um),
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
}

type bookReviewListReview struct {
	ID         int64               `json:"id"`         // レビューID
	Impression string              `json:"impression"` // 感想
	User       *bookReviewListUser `json:"user"`       // 投稿者
	CreatedAt  string              `json:"createdAt"`  // 登録日時
	UpdatedAt  string              `json:"updatedAt"`  // 更新日時
}

func newBookReviewListReview(r *entity.Review, u *entity.User) *bookReviewListReview {
	return &bookReviewListReview{
		ID:         r.Id,
		Impression: r.Impression,
		User:       newBookReviewListUser(u),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newBookReviewListReviews(rs entity.Reviews, um map[string]*entity.User) []*bookReviewListReview {
	res := make([]*bookReviewListReview, len(rs))
	for i, r := range rs {
		u := um[r.UserId]
		res[i] = newBookReviewListReview(r, u)
	}
	return res
}

type bookReviewListUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newBookReviewListUser(u *entity.User) *bookReviewListUser {
	if u == nil {
		return &bookReviewListUser{
			Username: "unknown",
		}
	}

	return &bookReviewListUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

// ユーザー毎のレビュー一覧
type UserReviewListResponse struct {
	Reviews []*userReviewListReview `json:"reviewsList"` // レビュー一覧
	Limit   int64                   `json:"limit"`       // 取得上限数
	Offset  int64                   `json:"offset"`      // 取得開始位置
	Total   int64                   `json:"total"`       // 検索一致数
}

func NewUserReviewListResponse(
	rs entity.Reviews, bm map[int64]*entity.Book, limit, offset, total int64,
) *UserReviewListResponse {
	return &UserReviewListResponse{
		Reviews: newUserReviewListReviews(rs, bm),
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
}

type userReviewListReview struct {
	ID         int64               `json:"id"`         // レビューID
	Impression string              `json:"impression"` // 感想
	Book       *userReviewListBook `json:"book"`       // 書籍情報
	CreatedAt  string              `json:"createdAt"`  // 登録日時
	UpdatedAt  string              `json:"updatedAt"`  // 更新日時
}

func newUserReviewListReview(r *entity.Review, b *entity.Book) *userReviewListReview {
	return &userReviewListReview{
		ID:         r.Id,
		Impression: r.Impression,
		Book:       newUserReviewListBook(b),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newUserReviewListReviews(rs entity.Reviews, bm map[int64]*entity.Book) []*userReviewListReview {
	res := make([]*userReviewListReview, len(rs))
	for i, r := range rs {
		b := bm[r.BookId]
		res[i] = newUserReviewListReview(r, b)
	}
	return res
}

type userReviewListBook struct {
	ID           int64  `json:"id"`           // 書籍ID
	Title        string `json:"title"`        // タイトル
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newUserReviewListBook(b *entity.Book) *userReviewListBook {
	return &userReviewListBook{
		ID:           b.Id,
		Title:        b.Title,
		ThumbnailURL: b.ThumbnailUrl,
	}
}
