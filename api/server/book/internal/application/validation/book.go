package validation

import (
	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

// BookRequestValidation - Book関連のリクエストバリデータ
type BookRequestValidation interface {
	BookItem(in *input.BookItem) error
	CreateBookshelf(in *input.CreateBookshelf) error
	UpdateBookshelf(in *input.UpdateBookshelf) error
	CreateAndUpdateBooks(in *input.CreateAndUpdateBooks) error
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

func (v *bookRequestValidation) BookItem(in *input.BookItem) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to BookItem for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) CreateBookshelf(in *input.CreateBookshelf) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateBookshelf for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) UpdateBookshelf(in *input.UpdateBookshelf) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to UpdateBookshelf for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) CreateAndUpdateBooks(in *input.CreateAndUpdateBooks) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateAndUpdateBooks for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
