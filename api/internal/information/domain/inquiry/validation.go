//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/information/domain/$GOPACKAGE/$GOFILE
package inquiry

import "context"

type Validation interface {
	Inquiry(ctx context.Context, i *Inquiry) error
}
