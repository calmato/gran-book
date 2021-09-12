package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

// フォロー一覧
type FollowListResponse struct {
	Users  []*followListUser `json:"usersList"` // フォロー一覧
	Limit  int64             `json:"limit"`     // 取得上限数
	Offset int64             `json:"offset"`    // 取得開始位置
	Total  int64             `json:"total"`     // 検索一致数
}

func NewFollowListResponse(fs entity.Follows, limit, offset, total int64) *FollowListResponse {
	return &FollowListResponse{
		Users:  newFollowListUsers(fs),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

type followListUser struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
}

func newFollowListUser(f *entity.Follow) *followListUser {
	return &followListUser{
		ID:               f.Id,
		Username:         f.Username,
		ThumbnailURL:     f.ThumbnailUrl,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollow,
	}
}

func newFollowListUsers(fs entity.Follows) []*followListUser {
	res := make([]*followListUser, len(fs))
	for i := range fs {
		res[i] = newFollowListUser(fs[i])
	}
	return res
}

// フォロワー一覧
type FollowerListResponse struct {
	Users  []*followerListUser `json:"usersList"` // フォロワー一覧
	Limit  int64               `json:"limit"`     // 取得上限数
	Offset int64               `json:"offset"`    // 取得開始位置
	Total  int64               `json:"total"`     // 検索一致数
}

func NewFollowerListResponse(fs entity.Followers, limit, offset, total int64) *FollowerListResponse {
	return &FollowerListResponse{
		Users:  newFollowerListUsers(fs),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

type followerListUser struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
}

func newFollowerListUser(f *entity.Follower) *followerListUser {
	return &followerListUser{
		ID:               f.Id,
		Username:         f.Username,
		ThumbnailURL:     f.ThumbnailUrl,
		SelfIntroduction: f.SelfIntroduction,
		IsFollow:         f.IsFollow,
	}
}

func newFollowerListUsers(fs entity.Followers) []*followerListUser {
	res := make([]*followerListUser, len(fs))
	for i := range fs {
		res[i] = newFollowerListUser(fs[i])
	}
	return res
}

// プロフィール情報
type UserProfileResponse struct {
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
	Products         []*userProfileProduct `json:"productsList"`     // 出品商品一覧
}

func NewUserProfileResponse(p *entity.UserProfile) *UserProfileResponse {
	return &UserProfileResponse{
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

type userProfileProduct struct {
	ID           int64    `json:"id"`           // 商品ID
	Name         string   `json:"name"`         // 商品名
	ThumbnailURL string   `json:"thumbnailUrl"` // サムネイルURL
	Authors      []string `json:"authorsList"`  // 著者名一覧
}

// TODO: create
// func newUserProfileProduct() *userProfileProduct {
// 	return nil
// }

// TODO: create
func newUserProfileProducts() []*userProfileProduct {
	res := make([]*userProfileProduct, 0)
	return res
}
