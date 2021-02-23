package user

import "context"

// Uploader - Userアップローダ
type Uploader interface {
	Thumbnail(ctx context.Context, uid string, data []byte) (string, error)
}
