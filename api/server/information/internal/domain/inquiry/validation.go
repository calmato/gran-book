package inquiry

import "context"

// Validation - Bookドメインバリデーション
type Validation interface {
	Inquiry(ctx context.Context, i *Inquiry) error
}
