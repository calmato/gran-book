package registry

import (
	"github.com/calmato/gran-book/api/server/information/internal/application"
	"github.com/calmato/gran-book/api/server/information/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/information/internal/infrastructure/service"
	gcs "github.com/calmato/gran-book/api/server/information/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication application.AuthApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, s *gcs.Storage) *Registry {
	auth := authInjection()

	return &Registry{
		AuthApplication: auth,
	}
}

func authInjection() application.AuthApplication {
	ar := repository.NewAuthRepository()
	as := service.NewAuthService(ar)

	aa := application.NewAuthApplication(as)

	return aa
}
