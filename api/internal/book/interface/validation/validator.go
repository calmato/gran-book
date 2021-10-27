//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/book/interface/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/book"
)

var (
	errInvalidValidation = errors.New("validation: failed to convert request")
)

func toValidationError(field, message string) error {
	ve := &exception.ValidationError{
		Field:   field,
		Message: message,
	}

	return exception.ErrInvalidRequestValidation.New(errInvalidValidation, ve)
}

type BookRequestValidation interface {
	ListBookshelf(req *book.ListBookshelfRequest) error
	ListBookReview(req *book.ListBookReviewRequest) error
	ListUserReview(req *book.ListUserReviewRequest) error
	ListUserMonthlyResult(req *book.ListUserMonthlyResultRequest) error
	MultiGetBooks(req *book.MultiGetBooksRequest) error
	GetBook(req *book.GetBookRequest) error
	GetBookByIsbn(req *book.GetBookByIsbnRequest) error
	GetBookshelf(req *book.GetBookshelfRequest) error
	GetReview(req *book.GetReviewRequest) error
	CreateBook(req *book.CreateBookRequest) error
	UpdateBook(req *book.UpdateBookRequest) error
	ReadBookshelf(req *book.ReadBookshelfRequest) error
	ReadingBookshelf(req *book.ReadingBookshelfRequest) error
	StackedBookshelf(req *book.StackedBookshelfRequest) error
	WantBookshelf(req *book.WantBookshelfRequest) error
	ReleaseBookshelf(req *book.ReleaseBookshelfRequest) error
	DeleteBook(req *book.DeleteBookRequest) error
	DeleteBookshelf(req *book.DeleteBookshelfRequest) error
}
