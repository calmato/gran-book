package cmd

import (
	"github.com/calmato/gran-book/api/internal/book/application"
	"github.com/calmato/gran-book/api/internal/book/infrastructure/repository"
	dv "github.com/calmato/gran-book/api/internal/book/infrastructure/validation"
	"github.com/calmato/gran-book/api/internal/book/interface/server"
	rv "github.com/calmato/gran-book/api/internal/book/interface/validation"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/book"
)

type registry struct {
	book book.BookServiceServer
}

func newRegistry(db *database.Client) *registry {
	bookApplication, bookRequestValidation := bookInjector(db)

	return &registry{
		book: server.NewBookServer(bookRequestValidation, bookApplication),
	}
}

func bookInjector(db *database.Client) (application.BookApplication, rv.BookRequestValidation) {
	br := repository.NewBookRepository(db, datetime.Now)
	bdv := dv.NewBookDomainValidation(br)
	ba := application.NewBookApplication(bdv, br)

	brv := rv.NewBookRequestValidation()

	return ba, brv
}
