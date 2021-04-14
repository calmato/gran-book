package book

import "context"

// Repository - Bookリポジトリ
type Repository interface {
	ListAuthorByBookID(ctx context.Context, bookID int) ([]*Author, error)
	Show(ctx context.Context, bookID int) (*Book, error)
	ShowByIsbn(ctx context.Context, isbn string) (*Book, error)
	ShowBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Bookshelf, error)
	ShowOrCreateAuthor(ctx context.Context, a *Author) error
	Create(ctx context.Context, b *Book) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	Update(ctx context.Context, b *Book) error
	UpdateBookshelf(ctx context.Context, b *Bookshelf) error
	MultipleCreate(ctx context.Context, bs []*Book) error
	MultipleUpdate(ctx context.Context, bs []*Book) error
	Delete(ctx context.Context, bookID int) error
	DeleteBookshelf(ctx context.Context, bookshelfID int) error
	GetIDByIsbn(ctx context.Context, isbn string) (int, error)
	GetBookshelfIDByUserIDAndBookID(ctx context.Context, userID string, bookID int) (int, error)
}
