package v1

// フォロー一覧
type FollowListResponse struct {
	Users  []*FollowListUser `json:"usersList"` // フォロー一覧
	Limit  int64             `json:"limit"`     // 取得上限数
	Offset int64             `json:"offset"`    // 取得開始位置
	Total  int64             `json:"total"`     // 検索一致数
}

type FollowListUser struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
}

// フォロワー一覧
type FollowerListResponse struct {
	Users  []*FollowerListUser `json:"usersList"` // フォロワー一覧
	Limit  int64               `json:"limit"`     // 取得上限数
	Offset int64               `json:"offset"`    // 取得開始位置
	Total  int64               `json:"total"`     // 検索一致数
}

type FollowerListUser struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // 表示名
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	IsFollow         bool   `json:"isFollow"`         // 自分がフォローしているか
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
	Products         []*UserProfileProduct `json:"productsList"`     // 出品商品一覧
}

type UserProfileProduct struct {
	ID           int64    `json:"id"`           // 商品ID
	Name         string   `json:"name"`         // 商品名
	ThumbnailURL string   `json:"thumbnailUrl"` // サムネイルURL
	Authors      []string `json:"authorsList"`  // 著者名一覧
}
