package book

import (
	"time"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/book"
)

// Book - 書籍エンティティ
type Book struct {
	ID             int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	Title          string    `gorm:"default:null"`
	TitleKana      string    `gorm:"default:null"`
	Description    string    `gorm:"default:null"`
	Isbn           string    `gorm:"default:null"`
	Publisher      string    `gorm:"default:null"`
	PublishedOn    string    `gorm:"default:null"`
	ThumbnailURL   string    `gorm:"default:null"`
	RakutenURL     string    `gorm:"default:null"`
	RakutenSize    string    `gorm:"default:null"`
	RakutenGenreID string    `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"default:null;<-:create"`
	UpdatedAt      time.Time `gorm:"default:null"`
	Authors        Authors   `gorm:"many2many:authors_books"`
	Reviews        Reviews   `gorm:"foreignKey:BookID"`
}

type Books []*Book

func (b *Book) Proto() *book.Book {
	return &book.Book{
		Id:             int64(b.ID),
		Title:          b.Title,
		TitleKana:      b.TitleKana,
		Description:    b.Description,
		Isbn:           b.Isbn,
		Publisher:      b.Publisher,
		PublishedOn:    b.PublishedOn,
		ThumbnailUrl:   b.ThumbnailURL,
		RakutenUrl:     b.RakutenURL,
		RakutenSize:    b.RakutenSize,
		RakutenGenreId: b.RakutenGenreID,
		CreatedAt:      datetime.FormatTime(b.CreatedAt),
		UpdatedAt:      datetime.FormatTime(b.UpdatedAt),
		Authors:        b.Authors.Proto(),
	}
}

func (bs Books) Proto() []*book.Book {
	res := make([]*book.Book, len(bs))
	for i := range bs {
		res[i] = bs[i].Proto()
	}
	return res
}

type MonthlyResult struct {
	Year      int
	Month     int
	ReadTotal int
}

type MonthlyResults []*MonthlyResult

func (r *MonthlyResult) Proto() *book.MonthlyResult {
	return &book.MonthlyResult{
		Year:      int32(r.Year),
		Month:     int32(r.Month),
		ReadTotal: int64(r.ReadTotal),
	}
}

func (rs MonthlyResults) Proto() []*book.MonthlyResult {
	res := make([]*book.MonthlyResult, len(rs))
	for i := range rs {
		res[i] = rs[i].Proto()
	}
	return res
}

// Author - 著者エンティティ
type Author struct {
	ID        int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	Name      string    `gorm:"default:null"`
	NameKana  string    `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:null;<-:create"`
	UpdatedAt time.Time `gorm:"default:null"`
}

type Authors []*Author

func (a *Author) Proto() *book.Author {
	return &book.Author{
		Name:     a.Name,
		NameKana: a.NameKana,
	}
}

func (as Authors) Proto() []*book.Author {
	res := make([]*book.Author, len(as))
	for i := range as {
		res[i] = as[i].Proto()
	}
	return res
}

// Bookshelf - 本棚エンティティ
type Bookshelf struct {
	ID        int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	BookID    int       `gorm:"default:null"`
	UserID    string    `gorm:"default:null"`
	ReviewID  int       `gorm:"default:null;unique"`
	Status    int       `gorm:"default:null"`
	ReadOn    time.Time `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:null;<-:create"`
	UpdatedAt time.Time `gorm:"default:null"`
	Book      *Book     `gorm:"foreignKey:BookID"`
	Review    *Review   `gorm:"-"`
}

type Bookshelves []*Bookshelf

func (b *Bookshelf) Proto() *book.Bookshelf {
	return &book.Bookshelf{
		Id:        int64(b.ID),
		BookId:    int64(b.BookID),
		UserId:    b.UserID,
		ReviewId:  int64(b.ReviewID),
		Status:    book.BookshelfStatus(b.Status),
		ReadOn:    datetime.FormatDate(b.ReadOn),
		CreatedAt: datetime.FormatTime(b.CreatedAt),
		UpdatedAt: datetime.FormatTime(b.UpdatedAt),
	}
}

func (bs Bookshelves) Proto() []*book.Bookshelf {
	res := make([]*book.Bookshelf, len(bs))
	for i := range bs {
		res[i] = bs[i].Proto()
	}
	return res
}

// Review - レビューエンティティ
type Review struct {
	ID         int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	BookID     int       `gorm:"default:null"`
	UserID     string    `gorm:"default:null"`
	Score      int       `gorm:"default:null"`
	Impression string    `gorm:"default:null"`
	CreatedAt  time.Time `gorm:"default:null;<-:create"`
	UpdatedAt  time.Time `gorm:"default:null"`
}

type Reviews []*Review

func (r *Review) Proto() *book.Review {
	return &book.Review{
		Id:         int64(r.ID),
		BookId:     int64(r.BookID),
		UserId:     r.UserID,
		Score:      int32(r.Score),
		Impression: r.Impression,
		CreatedAt:  datetime.FormatTime(r.CreatedAt),
		UpdatedAt:  datetime.FormatTime(r.UpdatedAt),
	}
}

func (rs Reviews) Proto() []*book.Review {
	res := make([]*book.Review, len(rs))
	for i := range rs {
		res[i] = rs[i].Proto()
	}
	return res
}

// BookAuthor - 中間テーブル用
//nolint:revive
type BookAuthor struct {
	ID        int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	BookID    int       `gorm:"default:null"`
	AuthorID  int       `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:null;<-:create"`
	UpdatedAt time.Time `gorm:"default:null"`
}

// 本棚に保存している書籍のステータス
const (
	NoneStatus    int = iota // 未登録
	ReadStatus               // 読んだ本
	ReadingStatus            // 読んでいる本
	StackedStatus            // 積読本
	WantStatus               // 欲しい本
	ReleaseStatus            // 手放したい本
)
