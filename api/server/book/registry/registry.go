package registry

import (
	"github.com/calmato/gran-book/api/server/book/internal/application"
	"github.com/calmato/gran-book/api/server/book/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/book/internal/infrastructure/service"
	gcs "github.com/calmato/gran-book/api/server/book/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication application.AuthApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, s *gcs.Storage) *Registry {
	auth := authInjection(db, s)

	return &Registry{
		AuthApplication: auth,
	}
}

func authInjection(db *repository.Client, s *gcs.Storage) application.AuthApplication {
	ar := repository.NewAuthRepository()
	as := service.NewAuthService(ar)

	aa := application.NewAuthApplication(as)

	return aa
}
