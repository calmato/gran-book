package inquiry

import "time"

// Inquiry - お問い合わせエンティティ
type Inquiry struct {
	ID          int       `gorm:"primaryKey;autoIncrement;<-:create"`
	SenderId    string    `gorm:""`
	AdminId     string    `gorm:""`
	Subject     string    `gorm:""`
	Description string    `gorm:""`
	Email       string    `gorm:""`
	IsReplied   int       `gorm:""`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time `gorm:""`
}
