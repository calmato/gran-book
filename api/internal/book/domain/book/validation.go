//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/book/domain/$GOPACKAGE/$GOFILE
package book

import "context"

// Validation - Bookドメインバリデーション
type Validation interface {
	Book(ctx context.Context, b *Book) error
	Author(ctx context.Context, a *Author) error
	Bookshelf(ctx context.Context, bs *Bookshelf) error
	Review(ctx context.Context, rv *Review) error
}
