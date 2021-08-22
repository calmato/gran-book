package book

import "time"

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
	Authors        []*Author `gorm:"many2many:authors_books"`
	Reviews        []*Review `gorm:"foreignKey:BookID"`
}

// Author - 著者エンティティ
type Author struct {
	ID        int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	Name      string    `gorm:"default:null"`
	NameKana  string    `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:null;<-:create"`
	UpdatedAt time.Time `gorm:"default:null"`
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

// BookAuthor - 中間テーブル用
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
