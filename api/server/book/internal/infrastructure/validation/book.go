package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
)

type bookDomainValidation struct{}

// NewBookDomainValidation - Book関連のドメインバリデータ
func NewBookDomainValidation() book.Validation {
	return &bookDomainValidation{}
}

func (v *bookDomainValidation) Book(ctx context.Context, b *book.Book) error {
	return nil
}

func (v *bookDomainValidation) Author(ctx context.Context, a *book.Author) error {
	return nil
}

func (v *bookDomainValidation) Bookshelf(ctx context.Context, b *book.Bookshelf) error {
	return nil
}

func (v *bookDomainValidation) Category(ctx context.Context, b *book.Category) error {
	return nil
}

func (v *bookDomainValidation) Publisher(ctx context.Context, b *book.Publisher) error {
	return nil
}
