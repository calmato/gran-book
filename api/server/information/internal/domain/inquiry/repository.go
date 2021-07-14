package inquiry

import (
	"context"
)

//  Repository - Inquiryリポジトリ
type Repository interface {
	Create(ctx context.Context, i *Inquiry) error
}
