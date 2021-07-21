package inquiry

import "time"

// Inquiry - お問い合わせエンティティ
type Inquiry struct {
	ID          int       `gorm:"primaryKey;autoIncrement;<-:create"`
	SenderID    string    `gorm:""`
	AdminID     *string   `gorm:""`
	Subject     string    `gorm:""`
	Description string    `gorm:""`
	Email       string    `gorm:""`
	IsReplied   bool      `gorm:""`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time `gorm:""`
}
