package inquiry

import (
	"context"
)

// Service - Inquiryサービス
type Service interface {
	Create(ctx context.Context, i *Inquiry) error
}
