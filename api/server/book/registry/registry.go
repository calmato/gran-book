package registry

import (
	"github.com/calmato/gran-book/api/server/book/internal/application"
	rv "github.com/calmato/gran-book/api/server/book/internal/application/validation"
	"github.com/calmato/gran-book/api/server/book/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/book/internal/infrastructure/service"
	dv "github.com/calmato/gran-book/api/server/book/internal/infrastructure/validation"
	gcs "github.com/calmato/gran-book/api/server/book/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication application.AuthApplication
	BookApplication application.BookApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, s *gcs.Storage) *Registry {
	auth := authInjection()
	book := bookInjection(db)

	return &Registry{
		AuthApplication: auth,
		BookApplication: book,
	}
}

func authInjection() application.AuthApplication {
	ar := repository.NewAuthRepository()
	as := service.NewAuthService(ar)

	aa := application.NewAuthApplication(as)

	return aa
}

func bookInjection(db *repository.Client) application.BookApplication {
	br := repository.NewBookRepository(db)
	bdv := dv.NewBookDomainValidation()
	bs := service.NewBookService(bdv, br)
	brv := rv.NewBookRequestValidation()

	ba := application.NewBookApplication(brv, bs)

	return ba
}
