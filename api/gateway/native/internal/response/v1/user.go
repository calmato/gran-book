package v1

// フォロー一覧
type FollowListV1Response struct {
	Users  []*FollowListV1Response_User `json:"usersList,omitempty"` // フォロー一覧
	Limit  int64                        `json:"limit,omitempty"`     // 取得上限数
	Offset int64                        `json:"offset,omitempty"`    // 取得開始位置
	Total  int64                        `json:"total,omitempty"`     // 検索一致数
}

type FollowListV1Response_User struct {
	Id               string `json:"id,omitempty"`               // ユーザーID
	Username         string `json:"username,omitempty"`         // 表示名
	ThumbnailUrl     string `json:"thumbnailUrl,omitempty"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction,omitempty"` // 自己紹介
	IsFollow         bool   `json:"isFollow,omitempty"`         // 自分がフォローしているか
}

// フォロワー一覧
type FollowerListV1Response struct {
	Users  []*FollowerListV1Response_User `json:"usersList,omitempty"` // フォロワー一覧
	Limit  int64                          `json:"limit,omitempty"`     // 取得上限数
	Offset int64                          `json:"offset,omitempty"`    // 取得開始位置
	Total  int64                          `json:"total,omitempty"`     // 検索一致数
}

type FollowerListV1Response_User struct {
	Id               string `json:"id,omitempty"`               // ユーザーID
	Username         string `json:"username,omitempty"`         // 表示名
	ThumbnailUrl     string `json:"thumbnailUrl,omitempty"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction,omitempty"` // 自己紹介
	IsFollow         bool   `json:"isFollow,omitempty"`         // 自分がフォローしているか
}

// プロフィール情報
type UserProfileV1Response struct {
	Id               string                           `json:"id,omitempty"`               // ユーザーID
	Username         string                           `json:"username,omitempty"`         // 表示名
	ThumbnailUrl     string                           `json:"thumbnailUrl,omitempty"`     // サムネイルURL
	SelfIntroduction string                           `json:"selfIntroduction,omitempty"` // 自己紹介
	IsFollow         bool                             `json:"isFollow,omitempty"`         // 自分がフォローしているか
	IsFollower       bool                             `json:"isFollower,omitempty"`       // 自分がフォローされているか
	FollowCount      int64                            `json:"followCount,omitempty"`      // フォロー数
	FollowerCount    int64                            `json:"followerCount,omitempty"`    // フォロワー数
	Rating           int32                            `json:"rating,omitempty"`           // ユーザーからの平均評価
	ReviewCount      int64                            `json:"reviewCount,omitempty"`      // ユーザーからのレビュー数
	Products         []*UserProfileV1Response_Product `json:"productsList,omitempty"`     // 出品商品一覧
}

type UserProfileV1Response_Product struct {
	Id           int64    `json:"id,omitempty"`           // 商品ID
	Name         string   `json:"name,omitempty"`         // 商品名
	ThumbnailUrl string   `json:"thumbnailUrl,omitempty"` // サムネイルURL
	Authors      []string `json:"authorsList,omitempty"`  // 著者名一覧
}
