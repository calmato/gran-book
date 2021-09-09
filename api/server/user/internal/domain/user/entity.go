package user

import (
	"time"

	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
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
	IsFollowing      bool           `gorm:"-"`
	IsFollowed       bool           `gorm:"-"`
	FollowCount      int            `gorm:"-"`
	FollowerCount    int            `gorm:"-"`
}

type Users []*User

func (u *User) Proto() *pb.User {
	return &pb.User{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           pb.Gender(u.Gender),
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}

func (u *User) Auth() *pb.Auth {
	return &pb.Auth{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           pb.Gender(u.Gender),
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             pb.Role(u.Role),
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		PostalCode:       u.PostalCode,
		Prefecture:       u.Prefecture,
		City:             u.City,
		AddressLine1:     u.AddressLine1,
		AddressLine2:     u.AddressLine2,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}

func (u *User) Admin() *pb.Admin {
	return &pb.Admin{
		Id:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             pb.Role(u.Role),
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}

func (u *User) Profile() *pb.UserProfile {
	return &pb.UserProfile{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		IsFollow:         u.IsFollowing,
		IsFollower:       u.IsFollowed,
		FollowCount:      int64(u.FollowCount),
		FollowerCount:    int64(u.FollowerCount),
	}
}

func (us Users) Proto() []*pb.User {
	res := make([]*pb.User, len(us))
	for i := range us {
		res[i] = us[i].Proto()
	}
	return res
}

func (us Users) Admin() []*pb.Admin {
	res := make([]*pb.Admin, len(us))
	for i := range us {
		res[i] = us[i].Admin()
	}
	return res
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

type Follows []*Follow

func (f *Follow) Proto() *pb.Follow {
	return &pb.Follow{
		Id:               f.FollowID,
		Username:         f.Username,
		ThumbnailUrl:     f.ThumbnailURL,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollowing,
	}
}

func (fs Follows) Proto() []*pb.Follow {
	res := make([]*pb.Follow, len(fs))
	for i := range fs {
		res[i] = fs[i].Proto()
	}
	return res
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

type Followers []*Follower

func (f *Follower) Proto() *pb.Follower {
	return &pb.Follower{
		Id:               f.FollowerID,
		Username:         f.Username,
		ThumbnailUrl:     f.ThumbnailURL,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollowing,
	}
}

func (fs Followers) Proto() []*pb.Follower {
	res := make([]*pb.Follower, len(fs))
	for i := range fs {
		res[i] = fs[i].Proto()
	}
	return res
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
