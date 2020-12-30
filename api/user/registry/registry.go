package registry

import (
	"github.com/calmato/gran-book/api/user/internal/application"
	"github.com/calmato/gran-book/api/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/user/lib/firebase/authentication"
)

// Registry - DIコンテナ
type Registry struct {
	UserApplication application.UserApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, auth *authentication.Auth) *Registry {
	ua := userInjection(db, auth)

	return &Registry{
		UserApplication: ua,
	}
}

func userInjection(db *repository.Client, auth *authentication.Auth) application.UserApplication {
	ur := repository.NewUserRepository(db, auth)
	us := service.NewUserService(ur)
	ua := application.NewUserApplication(us)

	return ua
}
