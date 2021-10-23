package v2

import (
	"strings"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
)

// 本棚の書籍情報
type BookshelfResponse struct {
	ID           int64               `json:"id"`           // 書籍ID
	Title        string              `json:"title"`        // タイトル
	TitleKana    string              `json:"titleKana"`    // タイトル(かな)
	Description  string              `json:"description"`  // 説明
	Isbn         string              `json:"isbn"`         // ISBN
	Publisher    string              `json:"publisher"`    // 出版社名
	PublishedOn  string              `json:"publishedOn"`  // 出版日
	ThumbnailURL string              `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string              `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string              `json:"size"`         // 楽天書籍サイズ
	Author       string              `json:"author"`       // 著者名一覧
	AuthorKana   string              `json:"author_kana"`  // 著者名一覧(かな)
	Bookshelf    *bookshelfBookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	Reviews      []*bookshelfReview  `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64               `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64               `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64               `json:"reviewTotal"`  // レビュー検索一致件
	CreatedAt    string              `json:"createdAt"`    // 登録日時
	UpdatedAt    string              `json:"updatedAt"`    // 更新日時
}

func NewBookshelfResponse(
	bs *entity.Bookshelf,
	b *entity.Book,
	rs entity.Reviews,
	um map[string]*entity.User,
	reviewLimit, reviewOffset, reviewTotal int64,
) *BookshelfResponse {
	return &BookshelfResponse{
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
		Bookshelf:    newBookshelfBookshelf(bs),
		Reviews:      newBookshelfReviews(rs, um),
		ReviewLimit:  reviewLimit,
		ReviewOffset: reviewOffset,
		ReviewTotal:  reviewTotal,
	}
}

type bookshelfBookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewID  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

func newBookshelfBookshelf(bs *entity.Bookshelf) *bookshelfBookshelf {
	return &bookshelfBookshelf{
		Status:    bs.Status().Name(),
		ReadOn:    bs.ReadOn,
		ReviewID:  bs.ReviewId,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: bs.UpdatedAt,
	}
}

type bookshelfReview struct {
	ID         int64          `json:"id"`         // レビューID
	Impression string         `json:"impression"` // 感想
	User       *bookshelfUser `json:"user"`       // 投稿者
	CreatedAt  string         `json:"createdAt"`  // 登録日時
	UpdatedAt  string         `json:"updatedAt"`  // 更新日時
}

func newBookshelfReview(r *entity.Review, u *entity.User) *bookshelfReview {
	return &bookshelfReview{
		ID:         r.Id,
		Impression: r.Impression,
		User:       newBookshelfUser(u),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func newBookshelfReviews(rs entity.Reviews, um map[string]*entity.User) []*bookshelfReview {
	res := make([]*bookshelfReview, len(rs))
	for i, r := range rs {
		u := um[r.UserId]
		res[i] = newBookshelfReview(r, u)
	}
	return res
}

type bookshelfUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newBookshelfUser(u *entity.User) *bookshelfUser {
	if u == nil {
		return &bookshelfUser{
			Username: "unknown",
		}
	}

	return &bookshelfUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

// 本棚の書籍一覧
type BookshelfListResponse struct {
	Books  []*bookshelfListBook `json:"booksList"` // 書籍一覧
	Limit  int64                `json:"limit"`     // 取得上限数
	Offset int64                `json:"offset"`    // 取得開始位置
	Total  int64                `json:"total"`     // 検索一致数
}

func NewBookshelfListResponse(
	bss entity.Bookshelves, bm map[int64]*entity.Book, limit, offset, total int64,
) *BookshelfListResponse {
	return &BookshelfListResponse{
		Books:  newBookshelfListBooks(bss, bm),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

type bookshelfListBook struct {
	ID           int64                   `json:"id"`           // 書籍ID
	Title        string                  `json:"title"`        // タイトル
	TitleKana    string                  `json:"titleKana"`    // タイトル(かな)
	Description  string                  `json:"description"`  // 説明
	Isbn         string                  `json:"isbn"`         // ISBN
	Publisher    string                  `json:"publisher"`    // 出版社名
	PublishedOn  string                  `json:"publishedOn"`  // 出版日
	ThumbnailURL string                  `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string                  `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string                  `json:"size"`         // 楽天書籍サイズ
	Author       string                  `json:"author"`       // 著者名一覧
	AuthorKana   string                  `json:"authorKana"`   // 著者名一覧(かな)
	Bookshelf    *bookshelfListBookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string                  `json:"createdAt"`    // 登録日時
	UpdatedAt    string                  `json:"updatedAt"`    // 更新日時
}

func newBookshelfListBook(bs *entity.Bookshelf, b *entity.Book) *bookshelfListBook {
	return &bookshelfListBook{
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
		Bookshelf:    newBookshelfListBookshelf(bs),
	}
}

func newBookshelfListBooks(bss entity.Bookshelves, bm map[int64]*entity.Book) []*bookshelfListBook {
	res := make([]*bookshelfListBook, 0, len(bss))
	for _, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}

		res = append(res, newBookshelfListBook(bs, b))
	}
	return res
}

type bookshelfListBookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewID  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

func newBookshelfListBookshelf(bs *entity.Bookshelf) *bookshelfListBookshelf {
	return &bookshelfListBookshelf{
		Status:    bs.Status().Name(),
		ReadOn:    bs.ReadOn,
		ReviewID:  bs.ReviewId,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: bs.UpdatedAt,
	}
}
