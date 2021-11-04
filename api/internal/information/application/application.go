//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/information/$GOPACKAGE/$GOFILE
package application

import (
	"context"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
)

type InquiryApplication interface {
	Create(ctx context.Context, i *inquiry.Inquiry) error
}
