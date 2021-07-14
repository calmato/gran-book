package inquiry

import "context"

// Validation - Inquiryドメインバリデーション
type Validation interface {
	Inquiry(ctx context.Context, i *Inquiry) error
}
