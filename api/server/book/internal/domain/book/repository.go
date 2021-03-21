package book

import "context"

// Repository - Bookリポジトリ
type Repository interface {
	ShowByIsbn(ctx context.Context, isbn string) (*Book, error)
	ShowPublisherByBookID(ctx context.Context, bookID int) (*Publisher, error)
	ShowAuthorsByBookID(ctx context.Context, bookID int) ([]*Author, error)
	ShowCategoriesByBookID(ctx context.Context, bookID int) ([]*Category, error)
	Create(ctx context.Context, b *Book) error
	CreateAuthor(ctx context.Context, a *Author) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateCategory(ctx context.Context, c *Category) error
	CreatePublisher(ctx context.Context, p *Publisher) error
	Update(ctx context.Context, b *Book) error
	MultipleCreate(ctx context.Context, bs []*Book) error
	MultipleUpdate(ctx context.Context, bs []*Book) error
}
