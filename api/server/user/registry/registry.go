package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/application"
	rv "github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/server/user/internal/infrastructure/validation"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/authentication"
	gcs "github.com/calmato/gran-book/api/server/user/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication application.AuthApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, auth *authentication.Auth, s *gcs.Storage) *Registry {
	aa := authInjection(db, auth, s)

	return &Registry{
		AuthApplication: aa,
	}
}

func authInjection(db *repository.Client, auth *authentication.Auth, s *gcs.Storage) application.AuthApplication {
	ur := repository.NewUserRepository(db, auth)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	arv := rv.NewAuthRequestValidation()
	aa := application.NewAuthApplication(arv, us)

	return aa
}
