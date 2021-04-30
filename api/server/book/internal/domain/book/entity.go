package book

import "time"

// Book - 書籍エンティティ
type Book struct {
	ID             int        `gorm:"primaryKey;autoIncrement;<-:create"`
	Title          string     `gorm:"uniqueIndex:ui_books_01;uniqueIndex:ui_books_02"`
	TitleKana      string     `gorm:""`
	Description    string     `gorm:""`
	Isbn           string     `gorm:""`
	Publisher      string     `gorm:""`
	PublishedOn    string     `gorm:""`
	ThumbnailURL   string     `gorm:""`
	RakutenURL     string     `gorm:""`
	RakutenGenreID string     `gorm:""`
	CreatedAt      time.Time  `gorm:"<-:create"`
	UpdatedAt      time.Time  `gorm:""`
	Authors        []*Author  `gorm:"many2many:authors_books"`
	Bookshelf      *Bookshelf `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
	Reviews        []*Review  `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}

// Author - 著者エンティティ
type Author struct {
	ID        int       `gorm:"primaryKey;autoIncrement;<-:create"`
	Name      string    `gorm:""`
	NameKana  string    `gorm:""`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

// Bookshelf - 本棚エンティティ
type Bookshelf struct {
	ID        int       `gorm:"primaryKey;autoIncrement;<-:create"`
	BookID    int       `gorm:"uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	UserID    string    `gorm:"uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	Status    int       `gorm:""`
	ReadOn    time.Time `gorm:""`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
	Book      *Book     `gorm:""`
}

// Review - レビューエンティティ
type Review struct {
	ID         int       `gorm:"primaryKey;autoIncrement;<-:create"`
	BookID     int       `gorm:"uniqueIndex:ui_reviews_01;uniqueIndex:ui_reviews_02"`
	UserID     string    `gorm:"uniqueIndex:ui_reviews_01;uniqueIndex:ui_reviews_02"`
	Score      int       `gorm:""`
	Impression string    `gorm:""`
	CreatedAt  time.Time `gorm:"<-:create"`
	UpdatedAt  time.Time `gorm:""`
}

// BookAuthor - 中間テーブル用
type BookAuthor struct {
	ID        int       `gorm:"primaryKey;autoIncrement;<-:create"`
	BookID    int       `gorm:"uniqueIndex:ui_authors_books_01;uniqueIndex:ui_authors_books_02"`
	AuthorID  int       `gorm:"uniqueIndex:ui_authors_books_01;uniqueIndex:ui_authors_books_02"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
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
