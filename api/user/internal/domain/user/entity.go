package user

import "time"

// User - Userエンティティ
type User struct {
	ID               int64  `gorm:"primaryKey;not null"`
	Username         string `gorm:"size:32;not null"`
	Gender           int32  `gorm:"not null;default:0"`
	Email            string `gorm:"size:256;not null;unique"`
	PasswordDigest   string `gorm:"size:256"`
	ThumbnailURL     string
	SelfIntroduction string `gorm:"size:256"`
	Lastname         string `gorm:"size:16"`
	Firstname        string `gorm:"size:16"`
	LastnameKana     string `gorm:"size:32"`
	FirstnameKana    string `gorm:"size:32"`
	PostalCode       string `gorm:"size:8"`
	Prefecture       string `gorm:"size:32"`
	City             string `gorm:"size:32"`
	AddressLine1     string `gorm:"size:64"`
	AddressLine2     string `gorm:"size:64"`
	PhoneNumber      string `gorm:"size:16"`
	Role             int32  `gorm:"not null;default:0"`
	InstanceID       string
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}
