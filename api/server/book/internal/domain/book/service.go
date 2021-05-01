package book

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain"
)

// Service - Bookサービス
type Service interface {
	ListBookshelf(ctx context.Context, q *domain.ListQuery) ([]*Bookshelf, error)
	ListBookshelfCount(ctx context.Context, q *domain.ListQuery) (int, error)
	Show(ctx context.Context, bookID int) (*Book, error)
	ShowByIsbn(ctx context.Context, isbn string) (*Book, error)
	ShowBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Bookshelf, error)
	ShowReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Review, error)
	Create(ctx context.Context, b *Book) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	Update(ctx context.Context, b *Book) error
	UpdateBookshelf(ctx context.Context, b *Bookshelf) error
	MultipleCreate(ctx context.Context, bs []*Book) error
	MultipleUpdate(ctx context.Context, bs []*Book) error
	Delete(ctx context.Context, bookID int) error
	DeleteBookshelf(ctx context.Context, bookshelfID int) error
	Validation(ctx context.Context, b *Book) error
	ValidationAuthor(ctx context.Context, a *Author) error
	ValidationBookshelf(ctx context.Context, b *Bookshelf) error
}
