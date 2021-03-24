package validation

import (
	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

// BookRequestValidation - Book関連のリクエストバリデータ
type BookRequestValidation interface {
	BookItem(in *input.BookItem) error
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

func (v *bookRequestValidation) CreateAndUpdateBooks(in *input.CreateAndUpdateBooks) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to CreateAndUpdateBooks for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
