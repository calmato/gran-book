package user

import "time"

// User - Userエンティティ
type User struct {
	ID               string    `gorm:"primaryKey;not null;<-:create"`
	Username         string    `gorm:"size:32;not null"`
	Gender           int32     `gorm:"not null;default:0"`
	Email            string    `gorm:"size:256"`
	PhoneNumber      string    `gorm:"size:16"`
	Role             int32     `gorm:"not null;default:0"`
	Password         string    `gorm:"-"`
	ThumbnailURL     string    `gorm:""`
	SelfIntroduction string    `gorm:"size:256"`
	LastName         string    `gorm:"size:16"`
	FirstName        string    `gorm:"size:16"`
	LastNameKana     string    `gorm:"size:32"`
	FirstNameKana    string    `gorm:"size:32"`
	PostalCode       string    `gorm:"size:16"`
	Prefecture       string    `gorm:"size:32"`
	City             string    `gorm:"size:32"`
	AddressLine1     string    `gorm:"size:64"`
	AddressLine2     string    `gorm:"size:64"`
	InstanceID       string    `gorm:"size:256"`
	CreatedAt        time.Time `gorm:"not null;<-:create"`
	UpdatedAt        time.Time `gorm:"not null"`
}
