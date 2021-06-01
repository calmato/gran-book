package notification

import "time"

// Notification - お知らせエンティティ
type Notification struct {
	ID					int			`gorm:primaryKey;autoIncrement;<-:create"`
	AuthorID		string
	EditorID		string
	CategoryID	string
	Title				string
	Description	string
	Importance	string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

// Category	-	カテゴリーエンティティ
type Category struct {
	ID					int	`gorm:primaryKey;autoIncrement;<-:create"`
	Name				string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}
