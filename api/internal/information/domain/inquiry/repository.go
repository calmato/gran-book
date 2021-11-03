//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/information/domain/$GOPACKAGE/$GOFILE
package inquiry

import "context"

// Repository - Inquirtyリポジトリ
type Repository interface {
	Create(ctx context.Context, i *Inquiry) error
}
