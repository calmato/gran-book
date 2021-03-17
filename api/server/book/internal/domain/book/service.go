package book

import "context"

// Service - Bookサービス
type Service interface {
	Create(ctx context.Context, b *Book) error
	CreateAuthor(ctx context.Context, b *Author) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateCategory(ctx context.Context, b *Category) error
	CreatePublisher(ctx context.Context, b *Publisher) error
}
