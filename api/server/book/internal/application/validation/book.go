package validation

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

// BookRequestValidation - Book関連のリクエストバリデータ
type BookRequestValidation interface {
	Book(in *input.Book) error
	Bookshelf(in *input.Bookshelf) error
	ListBookByBookIDs(in *input.ListBookByBookIDs) error
	ListBookshelf(in *input.ListBookshelf) error
	ListBookReview(in *input.ListBookReview) error
	ListUserReview(in *input.ListUserReview) error
}

type bookRequestValidation struct {
	validator RequestValidator
}

// NewBookRequestValidation - BookRequestValidationの生成
func NewBookRequestValidation() BookRequestValidation {
	rv := NewRequestValidator()

	return &bookRequestValidation{
		validator: rv,
	}
}

func (v *bookRequestValidation) Book(in *input.Book) error {
	ves := v.validator.Run(in, "")

	for i, a := range in.Authors {
		prefix := fmt.Sprintf("authors[%d].", i)
		ves = append(ves, v.validator.Run(a, prefix)...)
	}

	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to Book for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) Bookshelf(in *input.Bookshelf) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to Bookshelf for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) ListBookByBookIDs(in *input.ListBookByBookIDs) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListBookByBookIDs for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) ListBookshelf(in *input.ListBookshelf) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListBookshelf for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) ListBookReview(in *input.ListBookReview) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListBookReview for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) ListUserReview(in *input.ListUserReview) error {
	ves := v.validator.Run(in, "")
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to ListUserReview for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
