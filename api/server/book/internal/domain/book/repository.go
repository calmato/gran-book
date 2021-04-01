package book

import "context"

// Repository - Bookリポジトリ
type Repository interface {
	Show(ctx context.Context, bookID int) (*Book, error)
	ShowByIsbn(ctx context.Context, isbn string) (*Book, error)
	ShowAuthorsByBookID(ctx context.Context, bookID int) ([]*Author, error)
	ShowBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Bookshelf, error)
	ShowCategoriesByBookID(ctx context.Context, bookID int) ([]*Category, error)
	Create(ctx context.Context, b *Book) error
	CreateAuthor(ctx context.Context, a *Author) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateCategory(ctx context.Context, c *Category) error
	Update(ctx context.Context, b *Book) error
	UpdateBookshelf(ctx context.Context, b *Bookshelf) error
	MultipleCreate(ctx context.Context, bs []*Book) error
	MultipleUpdate(ctx context.Context, bs []*Book) error
}
