package user

import (
	"time"

	"gorm.io/gorm"
)

// User - Userエンティティ
type User struct {
	ID               string         `gorm:"default:null;primaryKey;<-:create"`
	Username         string         `gorm:"default:null"`
	Gender           int            `gorm:"default:null"`
	Email            string         `gorm:"default:null"`
	PhoneNumber      string         `gorm:"default:null"`
	Role             int            `gorm:"default:null"`
	Password         string         `gorm:"-"`
	ThumbnailURL     string         `gorm:"default:null"`
	SelfIntroduction string         `gorm:"default:null"`
	LastName         string         `gorm:"default:null"`
	FirstName        string         `gorm:"default:null"`
	LastNameKana     string         `gorm:"default:null"`
	FirstNameKana    string         `gorm:"default:null"`
	PostalCode       string         `gorm:"default:null"`
	Prefecture       string         `gorm:"default:null"`
	City             string         `gorm:"default:null"`
	AddressLine1     string         `gorm:"default:null"`
	AddressLine2     string         `gorm:"default:null"`
	InstanceID       string         `gorm:"default:null"`
	CreatedAt        time.Time      `gorm:"default:null;<-:create"`
	UpdatedAt        time.Time      `gorm:"default:null"`
	DeletedAt        gorm.DeletedAt `gorm:"default:null"`
	Follows          []*Follow      `gorm:"foreignKey:FollowID"`
	Followers        []*Follower    `gorm:"foreignKey:FollowerID"`
}

// Relationship - Relationshipエンティティ (中間テーブル)
type Relationship struct {
	ID         int       `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	FollowID   string    `gorm:"default:null"`
	FollowerID string    `gorm:"default:null"`
	CreatedAt  time.Time `gorm:"default:null"`
	UpdatedAt  time.Time `gorm:"default:null"`
}

// Follow - フォローしているUserのエンティティ
type Follow struct {
	FollowID         string `gorm:"<-:false"`
	Username         string `gorm:"<-:false"`
	ThumbnailURL     string `gorm:"<-:false"`
	SelfIntroduction string `gorm:"<-:false"`
	IsFollowing      bool   `gorm:"-"`
	IsFollowed       bool   `gorm:"-"`
	FollowCount      int    `gorm:"-"`
	FollowerCount    int    `gorm:"-"`
}

// Follower - フォローされているUserのエンティティ
type Follower struct {
	FollowerID       string `gorm:"<-:false"`
	Username         string `gorm:"<-:false"`
	ThumbnailURL     string `gorm:"<-:false"`
	SelfIntroduction string `gorm:"<-:false"`
	IsFollowing      bool   `gorm:"-"`
	IsFollowed       bool   `gorm:"-"`
	FollowCount      int    `gorm:"-"`
	FollowerCount    int    `gorm:"-"`
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
