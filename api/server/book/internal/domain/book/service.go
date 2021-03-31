package book

import "context"

// Service - Bookサービス
type Service interface {
	ShowByIsbn(ctx context.Context, isbn string) (*Book, error)
	Create(ctx context.Context, b *Book) error
	CreateAuthor(ctx context.Context, b *Author) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateCategory(ctx context.Context, b *Category) error
	Update(ctx context.Context, b *Book) error
	MultipleCreate(ctx context.Context, bs []*Book) error
	MultipleUpdate(ctx context.Context, bs []*Book) error
	Validation(ctx context.Context, b *Book) error
	ValidationAuthor(ctx context.Context, a *Author) error
	ValidationCategory(ctx context.Context, c *Category) error
}
