package user

import "time"

// User - Userエンティティ
type User struct {
	ID               string `gorm:"primaryKey;not null;<-:create"`
	Username         string `gorm:"size:32;not null"`
	Email            string `gorm:"size:256"`
	Password         string `gorm:"-"`
	Gender           int32  `gorm:"not null;default:0"`
	Role             int32  `gorm:"not null;default:0"`
	ThumbnailURL     string
	SelfIntroduction string    `gorm:"size:256"`
	Lastname         string    `gorm:"size:16"`
	Firstname        string    `gorm:"size:16"`
	LastnameKana     string    `gorm:"size:32"`
	FirstnameKana    string    `gorm:"size:32"`
	PostalCode       string    `gorm:"size:16"`
	Prefecture       string    `gorm:"size:32"`
	City             string    `gorm:"size:32"`
	AddressLine1     string    `gorm:"size:64"`
	AddressLine2     string    `gorm:"size:64"`
	PhoneNumber      string    `gorm:"size:16"`
	InstanceID       string    `gorm:"size:256"`
	CreatedAt        time.Time `gorm:"not null;<-:create"`
	UpdatedAt        time.Time `gorm:"not null"`
}
