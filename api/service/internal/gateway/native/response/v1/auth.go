package v1

import "github.com/calmato/gran-book/api/service/internal/gateway/native/entity"

// 認証情報
type AuthResponse struct {
	*entity.Auth
}

// サムネイルURL
type AuthThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
