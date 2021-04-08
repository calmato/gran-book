package book

import "time"

// Book - 書籍エンティティ
type Book struct {
	ID             int        `gorm:"primaryKey;not null;<-:create"`
	Title          string     `gorm:"size:64;not null;uniqueIndex:ui_books_01;uniqueIndex:ui_books_02"`
	TitleKana      string     `gorm:"size:128"`
	Description    string     `gorm:"size:2000"`
	Isbn           string     `gorm:"size:16;not null"`
	Publisher      string     `gorm:"size:64;not null"`
	PublishedOn    time.Time  `gorm:"not null"`
	ThumbnailURL   string     `gorm:""`
	RakutenURL     string     `gorm:""`
	RakutenGenreID string     `gorm:""`
	CreatedAt      time.Time  `gorm:"not null;<-:create"`
	UpdatedAt      time.Time  `gorm:"not null"`
	Authors        []*Author  `gorm:"many2many:authors_books"`
	Bookshelf      *Bookshelf `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
	Reviews        []*Review  `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}

// Author - 著者エンティティ
type Author struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	Name      string    `gorm:"size:32;not null;unique"`
	NameKana  string    `gorm:"size:64;not null"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Bookshelf - 本棚エンティティ
type Bookshelf struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	BookID    int       `gorm:"not null;uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	UserID    string    `gorm:"not null;uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	Status    int       `gorm:"not null;default:0"`
	ReadOn    time.Time `gorm:""`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Review - レビューエンティティ
type Review struct {
	ID         int       `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	BookID     int       `gorm:"not null;uniqueIndex:ui_reviews_01;uniqueIndex:ui_reviews_02"`
	UserID     string    `gorm:"uniqueIndex:ui_reviews_01;uniqueIndex:ui_reviews_02"`
	Score      int       `gorm:"not null"`
	Impression string    `gorm:"size:2000;default:''"`
	CreatedAt  time.Time `gorm:"not null;<-:create"`
	UpdatedAt  time.Time `gorm:"not null"`
}

// BookAuthor - 中間テーブル用
type BookAuthor struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	BookID    int       `gorm:"not null;uniqueIndex:ui_authors_books_01;uniqueIndex:ui_authors_books_02"`
	AuthorID  int       `gorm:"not null;uniqueIndex:ui_authors_books_01;uniqueIndex:ui_authors_books_02"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// 本棚に保存している書籍のステータス
const (
	NoneStatus    int = iota // 未登録
	ReadStatus               // 読んだ本
	ReadingStatus            // 読んでいる本
	StackStatus              // 積読本
	WantStatus               // 欲しい本
	ReleaseStatus            // 手放したい本
)
