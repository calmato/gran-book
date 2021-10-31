package book

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/pkg/database"
)

// Repository - Bookリポジトリ
type Repository interface {
	List(ctx context.Context, q *database.ListQuery) (Books, error)
	ListBookshelf(ctx context.Context, q *database.ListQuery) (Bookshelves, error)
	ListReview(ctx context.Context, q *database.ListQuery) (Reviews, error)
	Count(ctx context.Context, q *database.ListQuery) (int, error)
	CountBookshelf(ctx context.Context, q *database.ListQuery) (int, error)
	CountReview(ctx context.Context, q *database.ListQuery) (int, error)
	MultiGet(ctx context.Context, bookIDs []int) (Books, error)
	Get(ctx context.Context, bookID int) (*Book, error)
	GetByIsbn(ctx context.Context, isbn string) (*Book, error)
	GetBookIDByIsbn(ctx context.Context, isbn string) (int, error)
	GetBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Bookshelf, error)
	GetBookshelfIDByUserIDAndBookID(ctx context.Context, userID string, bookID int) (int, error)
	GetReview(ctx context.Context, reviewID int) (*Review, error)
	GetReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*Review, error)
	GetReviewIDByUserIDAndBookID(ctx context.Context, userID string, bookID int) (int, error)
	GetAuthorByName(ctx context.Context, name string) (*Author, error)
	GetAuthorIDByName(ctx context.Context, name string) (int, error)
	Create(ctx context.Context, b *Book) error
	CreateBookshelf(ctx context.Context, b *Bookshelf) error
	CreateReview(ctx context.Context, rv *Review) error
	CreateAuthor(ctx context.Context, a *Author) error
	Update(ctx context.Context, b *Book) error
	UpdateBookshelf(ctx context.Context, b *Bookshelf) error
	UpdateReview(ctx context.Context, rv *Review) error
	MultipleCreate(ctx context.Context, bs Books) error
	MultipleUpdate(ctx context.Context, bs Books) error
	Delete(ctx context.Context, bookID int) error
	DeleteBookshelf(ctx context.Context, bookshelfID int) error
}
