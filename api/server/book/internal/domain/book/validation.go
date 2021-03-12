package book

import "context"

// Validation - Bookドメインバリデーション
type Validation interface {
	Book(ctx context.Context, b *Book) error
	Author(ctx context.Context, b *Author) error
	Bookshelf(ctx context.Context, b *Bookshelf) error
	Category(ctx context.Context, b *Category) error
	Publisher(ctx context.Context, b *Publisher) error
}
