package validation

import (
	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

// BookRequestValidation - Book関連のリクエストバリデータ
type BookRequestValidation interface {
	Book(in *input.Book) error
	Bookshelf(in *input.Bookshelf) error
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
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to Book for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}

func (v *bookRequestValidation) Bookshelf(in *input.Bookshelf) error {
	ves := v.validator.Run(in)
	if len(ves) == 0 {
		return nil
	}

	err := xerrors.New("Failed to Bookshelf for RequestValidation")
	return exception.InvalidRequestValidation.New(err, ves...)
}
