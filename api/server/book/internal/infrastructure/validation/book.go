package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

type bookDomainValidation struct {
	bookRepository book.Repository
}

// NewBookDomainValidation - Book関連のドメインバリデータ
func NewBookDomainValidation(br book.Repository) book.Validation {
	return &bookDomainValidation{
		bookRepository: br,
	}
}

func (v *bookDomainValidation) Book(ctx context.Context, b *book.Book) error {
	err := v.uniqueCheckIsbn(ctx, b.ID, b.Isbn)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "isbn",
				Message: exception.CustomUniqueMessage,
			},
		}

		return exception.Conflict.New(err, ves...)
	}

	return nil
}

func (v *bookDomainValidation) Author(ctx context.Context, a *book.Author) error {
	return nil
}

func (v *bookDomainValidation) Bookshelf(ctx context.Context, b *book.Bookshelf) error {
	err := v.uniqueCheckBookshelf(ctx, b.ID, b.UserID, b.BookID)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "bookshelf",
				Message: exception.CustomUniqueMessage,
			},
		}

		return exception.Conflict.New(err, ves...)
	}

	return nil
}

func (v *bookDomainValidation) Review(ctx context.Context, rv *book.Review) error {
	err := v.uniqueCheckReview(ctx, rv.ID, rv.UserID, rv.BookID)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "review",
				Message: exception.CustomUniqueMessage,
			},
		}

		return exception.Conflict.New(err, ves...)
	}

	return nil
}

func (v *bookDomainValidation) uniqueCheckIsbn(ctx context.Context, id int, isbn string) error {
	if isbn == "" {
		return nil
	}

	bookID, _ := v.bookRepository.GetIDByIsbn(ctx, isbn)
	if bookID == 0 || bookID == id {
		return nil
	}

	return xerrors.New("This isbn is already exists.")
}

func (v *bookDomainValidation) uniqueCheckBookshelf(ctx context.Context, id int, userID string, bookID int) error {
	if userID == "" || bookID == 0 {
		return nil
	}

	bookshelfID, _ := v.bookRepository.GetBookshelfIDByUserIDAndBookID(ctx, userID, bookID)
	if bookshelfID == 0 || bookshelfID == id {
		return nil
	}

	return xerrors.New("This bookshelf is already exists.")
}

func (v *bookDomainValidation) uniqueCheckReview(ctx context.Context, id int, userID string, bookID int) error {
	if userID == "" || bookID == 0 {
		return nil
	}

	reviewID, _ := v.bookRepository.GetReviewIDByUserIDAndBookID(ctx, userID, bookID)
	if reviewID == 0 || reviewID == id {
		return nil
	}

	return xerrors.New("This review is already exists.")
}
