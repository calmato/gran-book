package user

import "context"

// Uploader - Userアップローダ
type Uploader interface {
	Thumbnail(ctx context.Context, userID string, data []byte) (string, error)
}
