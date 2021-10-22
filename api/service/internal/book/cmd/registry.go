package cmd

import (
	"github.com/calmato/gran-book/api/service/internal/book/application"
	"github.com/calmato/gran-book/api/service/internal/book/infrastructure/repository"
	dv "github.com/calmato/gran-book/api/service/internal/book/infrastructure/validation"
	"github.com/calmato/gran-book/api/service/internal/book/interface/server"
	rv "github.com/calmato/gran-book/api/service/internal/book/interface/validation"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/proto/book"
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
