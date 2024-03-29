package response

import "github.com/calmato/gran-book/api/internal/gateway/admin/v1/entity"

// 認証情報
type AuthResponse struct {
	*entity.Auth
}

// サムネイルURL
type AuthThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
