package v1

type FollowListResponse struct {
	Users  []*FollowListUser `json:"users"`
	Limit  int64             `json:"limit"`
	Offset int64             `json:"offset"`
	Total  int64             `json:"total"`
}

type FollowListUser struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	ThumbnailURL     string `json:"thumbnailUrl"`
	SelfIntroduction string `json:"selfIntroduction"`
	IsFollow         bool   `json:"isFollow"`
}

type FollowerListResponse struct {
	Users  []*FollowerListUser `json:"users"`
	Limit  int64               `json:"limit"`
	Offset int64               `json:"offset"`
	Total  int64               `json:"total"`
}

type FollowerListUser struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	ThumbnailURL     string `json:"thumbnailUrl"`
	SelfIntroduction string `json:"selfIntroduction"`
	IsFollow         bool   `json:"isFollow"`
}

type UserProfileResponse struct {
	ID               string                `json:"id"`
	Username         string                `json:"username"`
	ThumbnailURL     string                `json:"thumbnailUrl"`
	SelfIntroduction string                `json:"selfIntroduction"`
	IsFollow         bool                  `json:"isFollow"`
	IsFollower       bool                  `json:"isFollower"`
	FollowCount      int64                 `json:"followCount"`
	FollowerCount    int64                 `json:"followerCount"`
	Rating           int32                 `json:"rating"`
	ReviewCount      int64                 `json:"reviewCount"`
	Products         []*UserProfileProduct `json:"products"`
}

type UserProfileProduct struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	ThumbnailURL string   `json:"thumbnailUrl"`
	Authors      []string `json:"authros"`
}
