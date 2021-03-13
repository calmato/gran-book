package book

import "time"

// Book - 書籍エンティティ
type Book struct {
	ID           int64        `gorm:"primaryKey;not null;<-:create"`
	PublisherID  int64        `gorm:"not null;uniqueIndex:ui_books_01;uniqueIndex:ui_books_02"`
	Title        string       `gorm:"size:32;not null;uniqueIndex:ui_books_01;uniqueIndex:ui_books_02"`
	Description  string       `gorm:"size:1000;not null"`
	Isbn         string       `gorm:"size:16;not null"`
	ThumbnailURL string       `gorm:""`
	Version      string       `gorm:"size:64"`
	PublishedOn  time.Time    `gorm:"not null"`
	CreatedAt    time.Time    `gorm:"not null;<-:create"`
	UpdatedAt    time.Time    `gorm:"not null"`
	Publisher    *Publisher   `gorm:"foreignKey:PublisherID;constraint:OnDelete:SET NULL"`
	Bookshelfs   []*Bookshelf `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
	Authors      []*Author    `gorm:"many2many:authors_books"`
	Categories   []*Category  `gorm:"many2many:books_categories"`
}

// Publisher - 出版社エンティティ
type Publisher struct {
	ID        int64     `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	Name      string    `gorm:"size:32;not null;unique"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Author - 著者エンティティ
type Author struct {
	ID        int64     `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	Name      string    `gorm:"size:32;not null;unique"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Category - 書籍カテゴリエンティティ
type Category struct {
	ID        int64     `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	Name      string    `gorm:"size:32;not null;unique"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Bookshelf - 本棚エンティティ
type Bookshelf struct {
	ID         int64     `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	BookID     int64     `gorm:"not null;uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	UserID     string    `gorm:"not null;uniqueIndex:ui_bookshelfs_01;uniqueIndex:ui_bookshelfs_02"`
	Status     int32     `gorm:"not null;default:0"`
	Impression string    `gorm:"size:1000;default:''"`
	CreatedAt  time.Time `gorm:"not null;<-:create"`
	UpdatedAt  time.Time `gorm:"not null"`
}