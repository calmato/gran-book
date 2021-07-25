package chat

import "context"

// Uploader - Chatアップローダ
type Uploader interface {
	Image(ctx context.Context, roomID string, data []byte) (string, error)
}
