//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/book/$GOPACKAGE/$GOFILE
package application

import (
	"context"

	"github.com/calmato/gran-book/api/service/internal/book/domain/book"
	"github.com/calmato/gran-book/api/service/pkg/database"
)

type BookApplication interface {
	List(ctx context.Context, q *database.ListQuery) (book.Books, int, error)
	ListBookshelf(ctx context.Context, q *database.ListQuery) (book.Bookshelves, int, error)
	ListBookReview(ctx context.Context, bookID, limit, offset int) (book.Reviews, int, error)
	ListUserReview(ctx context.Context, userID string, limit, offset int) (book.Reviews, int, error)
	ListUserMonthlyResult(ctx context.Context, userID, since, until string) (book.MonthlyResults, error)
	MultiGet(ctx context.Context, bookIDs []int) (book.Books, error)
	Get(ctx context.Context, bookID int) (*book.Book, error)
	GetByIsbn(ctx context.Context, isbn string) (*book.Book, error)
	GetBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	GetBookshelfByUserIDAndBookIDWithRelated(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	GetReview(ctx context.Context, reviewID int) (*book.Review, error)
	GetReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Review, error)
	Create(ctx context.Context, b *book.Book) error
	CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	Update(ctx context.Context, b *book.Book) error
	UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	CreateOrUpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	MultipleCreate(ctx context.Context, bs book.Books) error
	MultipleUpdate(ctx context.Context, bs book.Books) error
	Delete(ctx context.Context, b *book.Book) error
	DeleteBookshelf(ctx context.Context, bs *book.Bookshelf) error
}
