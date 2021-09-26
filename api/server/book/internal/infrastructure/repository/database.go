package repository

import (
	"errors"

	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
)

const (
	authorTable     = "authors"
	authorBookTable = "authors_books"
	bookTable       = "books"
	bookshelfTable  = "bookshelves"
	reviewTable     = "reviews"
)

func toDBError(err error) error {
	err = database.ToDBError(err)
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, database.ErrInvalidTransaction):
		return exception.InvalidDatabaseTransaction.New(err)
	case errors.Is(err, database.ErrRecordNotFound):
		return exception.NotFound.New(err)
	case errors.Is(err, database.ErrUnknown):
		return exception.Unknown.New(err)
	default:
		return exception.ErrorInDatastore.New(err)
	}
}
