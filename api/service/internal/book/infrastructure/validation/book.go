package validation

import (
	"context"
	"errors"

	"github.com/calmato/gran-book/api/service/internal/book/domain/book"
	"github.com/calmato/gran-book/api/service/pkg/exception"
)

type bookDomainValidation struct {
	bookRepository book.Repository
}

var (
	errAlreadyExistsIsbn      = errors.New("validation: this isbn is already exists")
	errAlreadyExistsBookshelf = errors.New("validation: this bookshelf is already exists")
	errAlreadyExistsReview    = errors.New("validation: this review is already exists")
)

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

		return exception.ErrConflict.New(err, ves...)
	}

	return nil
}

func (v *bookDomainValidation) Author(ctx context.Context, a *book.Author) error {
	return nil
}

func (v *bookDomainValidation) Bookshelf(ctx context.Context, bs *book.Bookshelf) error {
	err := v.uniqueCheckBookshelf(ctx, bs.ID, bs.UserID, bs.BookID)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "bookshelf",
				Message: exception.CustomUniqueMessage,
			},
		}

		return exception.ErrConflict.New(err, ves...)
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

		return exception.ErrConflict.New(err, ves...)
	}

	return nil
}

func (v *bookDomainValidation) uniqueCheckIsbn(ctx context.Context, id int, isbn string) error {
	if isbn == "" {
		return nil
	}

	bookID, _ := v.bookRepository.GetBookIDByIsbn(ctx, isbn)
	if bookID == 0 || bookID == id {
		return nil
	}

	return errAlreadyExistsIsbn
}

func (v *bookDomainValidation) uniqueCheckBookshelf(ctx context.Context, id int, userID string, bookID int) error {
	if userID == "" || bookID == 0 {
		return nil
	}

	bookshelfID, _ := v.bookRepository.GetBookshelfIDByUserIDAndBookID(ctx, userID, bookID)
	if bookshelfID == 0 || bookshelfID == id {
		return nil
	}

	return errAlreadyExistsBookshelf
}

func (v *bookDomainValidation) uniqueCheckReview(ctx context.Context, id int, userID string, bookID int) error {
	if userID == "" || bookID == 0 {
		return nil
	}

	reviewID, _ := v.bookRepository.GetReviewIDByUserIDAndBookID(ctx, userID, bookID)
	if reviewID == 0 || reviewID == id {
		return nil
	}

	return errAlreadyExistsReview
}
