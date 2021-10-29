package entity

import (
	"strings"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
)

type BookOnBookshelf struct {
	ID           int64       `json:"id"`                     // 書籍ID
	Title        string      `json:"title"`                  // タイトル
	TitleKana    string      `json:"titleKana"`              // タイトル(かな)
	Description  string      `json:"description"`            // 説明
	Isbn         string      `json:"isbn"`                   // ISBN
	Publisher    string      `json:"publisher"`              // 出版社名
	PublishedOn  string      `json:"publishedOn"`            // 出版日
	ThumbnailURL string      `json:"thumbnailUrl"`           // サムネイルURL
	RakutenURL   string      `json:"rakutenUrl"`             // 楽天ショップURL
	Size         string      `json:"size"`                   // 楽天書籍サイズ
	Author       string      `json:"author"`                 // 著者名一覧
	AuthorKana   string      `json:"author_kana"`            // 著者名一覧(かな)
	Bookshelf    *Bookshelf  `json:"bookshelf"`              // ユーザーの本棚情報
	Reviews      BookReviews `json:"reviewsList,omitempty"`  // レビュー一覧
	ReviewLimit  int64       `json:"reviewLimit,omitempty"`  // レビュー取得上限
	ReviewOffset int64       `json:"reviewOffset,omitempty"` // レビュー取得開始位置
	ReviewTotal  int64       `json:"reviewTotal,omitempty"`  // レビュー検索一致件
	CreatedAt    string      `json:"createdAt"`              // 登録日時
	UpdatedAt    string      `json:"updatedAt"`              // 更新日時
}

type Bookshelf struct {
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	ReviewID  int64  `json:"reviewId"`  // レビューID
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

type BooksOnBookshelf []*BookOnBookshelf

func NewBookOnBookshelf(b *entity.Book, bs *entity.Bookshelf) *BookOnBookshelf {
	return &BookOnBookshelf{
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
		Bookshelf:    newBooshelf(bs),
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}

func newBooshelf(bs *entity.Bookshelf) *Bookshelf {
	return &Bookshelf{
		Status:    bs.Status().Name(),
		ReadOn:    bs.ReadOn,
		ReviewID:  bs.ReviewId,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: bs.UpdatedAt,
	}
}

func (b *BookOnBookshelf) Fill(rs BookReviews, limit, offset, total int64) {
	b.Reviews = rs
	b.ReviewLimit = limit
	b.ReviewOffset = offset
	b.ReviewTotal = total
}

func NewBooksOnBookshelf(bm map[int64]*entity.Book, bss entity.Bookshelves) BooksOnBookshelf {
	res := make(BooksOnBookshelf, 0, len(bss))
	for _, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}
		res = append(res, NewBookOnBookshelf(b, bs))
	}
	return res
}
