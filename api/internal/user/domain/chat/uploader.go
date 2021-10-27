//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package chat

import "context"

// Uploader - Chatアップローダ
type Uploader interface {
	Image(ctx context.Context, roomID string, data []byte) (string, error)
}
