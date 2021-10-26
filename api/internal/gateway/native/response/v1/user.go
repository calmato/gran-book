package v1

import "github.com/calmato/gran-book/api/internal/gateway/native/entity"

// フォロー一覧
type FollowListResponse struct {
	Users  entity.Follows `json:"usersList"` // フォロー一覧
	Limit  int64          `json:"limit"`     // 取得上限数
	Offset int64          `json:"offset"`    // 取得開始位置
	Total  int64          `json:"total"`     // 検索一致数
}

// フォロワー一覧
type FollowerListResponse struct {
	Users  entity.Followers `json:"usersList"` // フォロワー一覧
	Limit  int64            `json:"limit"`     // 取得上限数
	Offset int64            `json:"offset"`    // 取得開始位置
	Total  int64            `json:"total"`     // 検索一致数
}

// プロフィール情報
type UserProfileResponse struct {
	*entity.UserProfile
}
