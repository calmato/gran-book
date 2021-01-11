package registry

import (
	"github.com/calmato/gran-book/api/user/internal/application"
	rv "github.com/calmato/gran-book/api/user/internal/application/validation"
	"github.com/calmato/gran-book/api/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/user/internal/infrastructure/validation"
	"github.com/calmato/gran-book/api/user/lib/firebase/authentication"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication application.AuthApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, auth *authentication.Auth) *Registry {
	aa := authInjection(db, auth)

	return &Registry{
		AuthApplication: aa,
	}
}

func authInjection(db *repository.Client, auth *authentication.Auth) application.AuthApplication {
	ur := repository.NewUserRepository(db, auth)
	udv := validation.NewUserDomainValidation(ur)
	us := service.NewUserService(udv, ur, nil)

	arv := rv.NewAuthRequestValidation()
	aa := application.NewAuthApplication(arv, us)

	return aa
}
