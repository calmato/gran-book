package v1

import (
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
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
	CreatedAt    string              `json:"createdAt"`    // 登録日時
	UpdatedAt    string              `json:"updatedAt"`    // 更新日時
}

func NewBookshelfResponse(bs *entity.Bookshelf, b *entity.Book) *BookshelfResponse {
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
	}
}

type bookshelfBookshelf struct {
	ID         int64  `json:"id"`         // 本棚ID
	Status     string `json:"status"`     // 読書ステータス
	ReadOn     string `json:"readOn"`     // 読み終えた日
	Impression string `json:"impression"` // 感想
	CreatedAt  string `json:"createdAt"`  // 登録日時
	UpdatedAt  string `json:"updatedAt"`  // 更新日時
}

func newBookshelfBookshelf(bs *entity.Bookshelf) *bookshelfBookshelf {
	return &bookshelfBookshelf{
		ID:         bs.Id,
		Status:     bs.Status().Name(),
		ReadOn:     bs.ReadOn,
		Impression: "",
		CreatedAt:  bs.CreatedAt,
		UpdatedAt:  bs.UpdatedAt,
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
	for i, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}

		res[i] = newBookshelfListBook(bs, b)
	}
	return res
}

type bookshelfListBookshelf struct {
	ID        int64  `json:"id"`        // 本棚ID
	Status    string `json:"status"`    // 読書ステータス
	ReadOn    string `json:"readOn"`    // 読み終えた日
	CreatedAt string `json:"createdAt"` // 登録日時
	UpdatedAt string `json:"updatedAt"` // 更新日時
}

func newBookshelfListBookshelf(bs *entity.Bookshelf) *bookshelfListBookshelf {
	return &bookshelfListBookshelf{
		ID:        bs.Id,
		Status:    bs.Status().Name(),
		ReadOn:    bs.ReadOn,
		CreatedAt: bs.CreatedAt,
		UpdatedAt: bs.UpdatedAt,
	}
}
