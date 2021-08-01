package user

import (
	"time"

	"gorm.io/gorm"
)

// User - Userエンティティ
type User struct {
	ID               string    `gorm:"primaryKey;not null;<-:create"`
	Username         string    `gorm:"size:32;not null"`
	Gender           int       `gorm:"not null;default:0"`
	Email            string    `gorm:"size:256"`
	PhoneNumber      string    `gorm:"size:16"`
	Role             int       `gorm:"not null;default:0"`
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
	DeletedAt        gorm.DeletedAt
	Follows          []*Follow   `gorm:"foreignKey:FollowID;constraint:OnDelete:CASCADE"`
	Followers        []*Follower `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
}

// Relationship - Relationshipエンティティ (中間テーブル)
type Relationship struct {
	ID         int       `gorm:"primaryKey;not null;autoIncrement;<-:create"`
	FollowID   string    `gorm:"not null;uniqueIndex:ui_follows_01;uniqueIndex:ui_follows_02"`
	FollowerID string    `gorm:"not null;uniqueIndex:ui_follows_01;uniqueIndex:ui_follows_02"`
	CreatedAt  time.Time `gorm:"not null;<-:create"`
	UpdatedAt  time.Time `gorm:"not null"`
}

// Follow - フォローしているUserのエンティティ
type Follow struct {
	FollowID         string
	Username         string
	ThumbnailURL     string
	SelfIntroduction string
	IsFollowing      bool `gorm:"-"`
	IsFollowed       bool `gorm:"-"`
	FollowCount      int  `gorm:"-"`
	FollowerCount    int  `gorm:"-"`
}

// Follower - フォローされているUserのエンティティ
type Follower struct {
	FollowerID       string
	Username         string
	ThumbnailURL     string
	SelfIntroduction string
	IsFollowing      bool `gorm:"-"`
	IsFollowed       bool `gorm:"-"`
	FollowCount      int  `gorm:"-"`
	FollowerCount    int  `gorm:"-"`
}

// ユーザ権限
const (
	UserRole int = iota
	AdminRole
	DeveloperRole
	OperatorRole
)
