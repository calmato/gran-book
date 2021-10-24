package entity

import "github.com/calmato/gran-book/api/service/internal/gateway/entity"

type Follow struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
}

type Follows []*Follow

func NewFollow(f *entity.Follow) *Follow {
	return &Follow{
		ID:               f.Id,
		Username:         f.Username,
		ThumbnailURL:     f.ThumbnailUrl,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollow,
	}
}

func NewFollows(fs entity.Follows) Follows {
	res := make(Follows, len(fs))
	for i := range fs {
		res[i] = NewFollow(fs[i])
	}
	return res
}

type Follower struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
}

type Followers []*Follower

func NewFollower(f *entity.Follower) *Follower {
	return &Follower{
		ID:               f.Id,
		Username:         f.Username,
		ThumbnailURL:     f.ThumbnailUrl,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollow,
	}
}

func NewFollowers(fs entity.Followers) Followers {
	res := make(Followers, len(fs))
	for i := range fs {
		res[i] = NewFollower(fs[i])
	}
	return res
}

type UserProfile struct {
	ID               string                `json:"id"`               // ユーザーID
	Username         string                `json:"username"`         // 表示名
	ThumbnailURL     string                `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string                `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool                  `json:"isFollow"`         // 自分がフォローしているか
	IsFollower       bool                  `json:"isFollower"`       // 自分がフォローされているか
	FollowCount      int64                 `json:"followCount"`      // フォロー数
	FollowerCount    int64                 `json:"followerCount"`    // フォロワー数
	Rating           int32                 `json:"rating"`           // ユーザーからの平均評価
	ReviewCount      int64                 `json:"reviewCount"`      // ユーザーからのレビュー数
	Products         []*UserProfileProduct `json:"productsList"`     // 出品商品一覧
}

type UserProfileProduct struct {
	ID           int64    `json:"id"`           // 商品ID
	Name         string   `json:"name"`         // 商品名
	ThumbnailURL string   `json:"thumbnailUrl"` // サムネイルURL
	Authors      []string `json:"authorsList"`  // 著者名一覧
}

func NewUserProfile(p *entity.UserProfile) *UserProfile {
	return &UserProfile{
		ID:               p.Id,
		Username:         p.Username,
		ThumbnailURL:     p.ThumbnailUrl,
		SelfIntroduction: p.SelfIntroduction,
		IsFollow:         p.IsFollow,
		IsFollower:       p.IsFollower,
		FollowCount:      p.FollowCount,
		FollowerCount:    p.FollowerCount,
		Rating:           0,
		ReviewCount:      0,
		Products:         newUserProfileProducts(),
	}
}

// TODO: create
func newUserProfileProducts() []*UserProfileProduct {
	res := make([]*UserProfileProduct, 0)
	return res
}
