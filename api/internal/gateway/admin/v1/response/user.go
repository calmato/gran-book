package response

import "github.com/calmato/gran-book/api/internal/gateway/admin/entity"

type UserResponse struct {
	*entity.User
}

type UserListResponse struct {
	Users  entity.Users `json:"usersList"` // ユーザー一覧
	Limit  int64        `json:"limit"`     // 取得上限
	Offset int64        `json:"offset"`    // 取得開始位置
	Total  int64        `json:"total"`     // 検索一致数
}
