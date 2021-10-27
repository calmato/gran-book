package v1

import "github.com/calmato/gran-book/api/internal/gateway/admin/entity"

type AdminResponse struct {
	*entity.Admin
}

type AdminListResponse struct {
	Users  entity.Admins `json:"usersList"` // 管理者一覧
	Limit  int64         `json:"limit"`     // 取得上限
	Offset int64         `json:"offset"`    // 取得開始位置
	Total  int64         `json:"total"`     // 検索一致数
}

type AdminThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
