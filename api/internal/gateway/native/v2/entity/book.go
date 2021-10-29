package entity

import (
	"strings"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
)

type Book struct {
	ID           int64       `json:"id"`           // 書籍ID
	Title        string      `json:"title"`        // タイトル
	TitleKana    string      `json:"titleKana"`    // タイトル(かな)
	Description  string      `json:"description"`  // 説明
	Isbn         string      `json:"isbn"`         // ISBN
	Publisher    string      `json:"publisher"`    // 出版社名
	PublishedOn  string      `json:"publishedOn"`  // 出版日
	ThumbnailURL string      `json:"thumbnailUrl"` // サムネイルURL
	RakutenURL   string      `json:"rakutenUrl"`   // 楽天ショップURL
	Size         string      `json:"size"`         // 楽天書籍サイズ
	Author       string      `json:"author"`       // 著者名一覧
	AuthorKana   string      `json:"authorKana"`   // 著者名一覧(かな)
	Reviews      BookReviews `json:"reviewsList"`  // レビュー一覧
	ReviewLimit  int64       `json:"reviewLimit"`  // レビュー取得上限
	ReviewOffset int64       `json:"reviewOffset"` // レビュー取得開始位置
	ReviewTotal  int64       `json:"reviewTotal"`  // レビュー検索一致件数
	CreatedAt    string      `json:"createdAt"`    // 登録日時
	UpdatedAt    string      `json:"updatedAt"`    // 更新日時
}

func NewBook(b *entity.Book, rs BookReviews) *Book {
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
		Reviews:      rs,
	}
}

func (b *Book) Fill(limit, offset, total int64) {
	b.ReviewLimit = limit
	b.ReviewOffset = offset
	b.ReviewTotal = total
}
