package inquiry

import (
	"context"
)

// Service - Bookサービス
type Service interface {
	Create(ctx context.Context, i *Inquiry) error
}
