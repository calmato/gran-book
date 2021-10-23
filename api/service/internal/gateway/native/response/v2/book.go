package v2

import (
	"strings"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
)

// 書籍情報
type BookResponse struct {
	ID           int64         `json:"id"`           // 書籍ID
	Title        string        `json:"title"`        // タイトル
	TitleKana    string        `json:"titleKana"`    // タイトル(かな)
	Description  string        `json:"description"`  // 説明
	Isbn         string        `json:"isbn"`         // ISBN
	Publisher    string        `json:"publisher"`    // 出版社名
	PublishedOn  string        `json:"publishedOn"`  // 出版日
	ThumbnailURL string        `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string        `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string        `json:"size"`         // 楽天書籍サイズ
	Author       string        `json:"author"`       // 著者名一覧
	AuthorKana   string        `json:"authorKana"`   // 著者名一覧(かな)
	Reviews      []*bookReview `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64         `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64         `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64         `json:"reviewTotal"`  // レビュー検索一致件数
	CreatedAt    string        `json:"createdAt"`    // 登録日時
	UpdatedAt    string        `json:"updatedAt"`    // 更新日時
}

func NewBookResponse(
	b *entity.Book,
	rs entity.Reviews,
	um map[string]*entity.User,
	reviewLimit, reviewOffset, reviewTotal int64,
) *BookResponse {
	return &BookResponse{
		ID:           b.Id,
		Title:        b.Title,
		TitleKana:    b.TitleKana,
		Description:  b.Description,
		Isbn:         b.Isbn,
		Publisher:    b.Publisher,
		PublishedOn:  b.PublishedOn,
		ThumbnailURL: b.ThumbnailUrl,
		RakutenURL:   b.RakutenUrl,
		Size:         b.RakutenSize,
		Author:       strings.Join(b.AuthorNames(), "/"),
		AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
		Reviews:      newBookReviews(rs, um),
		ReviewLimit:  reviewLimit,
		ReviewOffset: reviewOffset,
		ReviewTotal:  reviewTotal,
	}
}

type bookReview struct {
	ID         int64     `json:"id"`         // レビューID
	Impression string    `json:"impression"` // 感想
	User       *bookUser `json:"user"`       // 投稿者
	CreatedAt  string    `json:"createdAt"`  // 登録日時
	UpdatedAt  string    `json:"updatedAt"`  // 更新日時
}

func newBookReview(r *entity.Review, u *entity.User) *bookReview {
	return &bookReview{
		ID:         r.Id,
		Impression: r.Impression,
		User:       newBookUser(u),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newBookReviews(rs entity.Reviews, um map[string]*entity.User) []*bookReview {
	res := make([]*bookReview, len(rs))
	for i, r := range rs {
		u := um[r.UserId]
		res[i] = newBookReview(r, u)
	}
	return res
}

type bookUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newBookUser(u *entity.User) *bookUser {
	if u == nil {
		return &bookUser{
			Username: "unknown",
		}
	}

	return &bookUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}
