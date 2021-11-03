package entity

import (
	"strings"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
)

type Book struct {
	ID           int64      `json:"id"`           // 書籍ID
	Title        string     `json:"title"`        // タイトル
	TitleKana    string     `json:"titleKana"`    // タイトル(かな)
	Description  string     `json:"description"`  // 説明
	Isbn         string     `json:"isbn"`         // ISBN
	Publisher    string     `json:"publisher"`    // 出版社名
	PublishedOn  string     `json:"published_on"` // 出版日
	ThumbnailURL string     `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string     `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string     `json:"size"`         // 楽天書籍サイズ
	Author       string     `json:"author"`       // 著者名一覧
	AuthorKana   string     `json:"authorKana"`   // 著者名一覧(かな)
	Bookshelf    *Bookshelf `json:"bookshelf"`    // ユーザーの本棚情報
	CreatedAt    string     `json:"createdAt"`    // 登録日時
	UpdatedAt    string     `json:"updatedAt"`    // 更新日時
}

type Books []*Book

func NewBook(b *entity.Book, bs *entity.Bookshelf) *Book {
	return &Book{
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
		Bookshelf:    NewBookshelf(bs),
	}
}

func NewBooks(bm map[int64]*entity.Book, bss entity.Bookshelves) Books {
	res := make(Books, 0, len(bss))
	for _, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}
	}
}
