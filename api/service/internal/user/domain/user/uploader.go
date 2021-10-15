//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package user

import "context"

// Uploader - Userアップローダ
type Uploader interface {
	Thumbnail(ctx context.Context, userID string, data []byte) (string, error)
}
