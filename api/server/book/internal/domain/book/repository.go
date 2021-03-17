package book

import "context"

// Repository - Bookリポジトリ
type Repository interface {
	Create(ctx context.Context, b *Book) error
	CreateAuthor(ctx context.Context, a *Author) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateCategory(ctx context.Context, c *Category) error
	CreatePublisher(ctx context.Context, p *Publisher) error
}
