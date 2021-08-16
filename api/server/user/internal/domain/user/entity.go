package user

import (
	"time"

	"gorm.io/gorm"
)

// User - Userエンティティ
type User struct {
	ID               string
	Username         string
	Gender           int
	Email            string
	PhoneNumber      string
	Role             int
	Password         string `gorm:"-"`
	ThumbnailURL     string
	SelfIntroduction string
	LastName         string
	FirstName        string
	LastNameKana     string
	FirstNameKana    string
	PostalCode       string
	Prefecture       string
	City             string
	AddressLine1     string
	AddressLine2     string
	InstanceID       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Follows          []*Follow   `gorm:"foreignKey:FollowID"`
	Followers        []*Follower `gorm:"foreignKey:FollowerID"`
}

// Relationship - Relationshipエンティティ (中間テーブル)
type Relationship struct {
	ID         int
	FollowID   string
	FollowerID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
	UserRole      int = iota // ユーザー
	AdminRole                // 管理者
	DeveloperRole            // 開発者
	OperatorRole             // 運用者
)

// ユーザ性別
const (
	UnkownGender        int = iota // 未選択
	MaleGender                     // 男性
	FemaleGender                   // 女性
	NotApplicableGender            // 適用不能
)
