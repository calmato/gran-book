package registry

import (
	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/infrastructure/repository"
	dv "github.com/calmato/gran-book/api/server/book/internal/infrastructure/validation"
	rv "github.com/calmato/gran-book/api/server/book/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
)

// Registry - DIコンテナ
type Registry struct {
	BookApplication       application.BookApplication
	BookRequestValidation rv.BookRequestValidation
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *database.Client) *Registry {
	bookApplication, bookRequestValidation := bookInjection(db)

	return &Registry{
		BookApplication:       bookApplication,
		BookRequestValidation: bookRequestValidation,
	}
}

func bookInjection(db *database.Client) (application.BookApplication, rv.BookRequestValidation) {
	br := repository.NewBookRepository(db)
	bdv := dv.NewBookDomainValidation(br)
	ba := application.NewBookApplication(bdv, br)

	brv := rv.NewBookRequestValidation()

	return ba, brv
}
