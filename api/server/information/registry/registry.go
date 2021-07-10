package registry

import (
	"github.com/calmato/gran-book/api/server/information/internal/application"
	rv "github.com/calmato/gran-book/api/server/information/internal/application/validation"
	"github.com/calmato/gran-book/api/server/information/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/information/internal/infrastructure/service"
	dv "github.com/calmato/gran-book/api/server/information/internal/infrastructure/validation"
	gcs "github.com/calmato/gran-book/api/server/information/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AuthApplication         application.AuthApplication
	NotificationApplication application.NotificationApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, s *gcs.Storage) *Registry {
	auth := authInjection()
	notification := notificationInjection(db)

	return &Registry{
		AuthApplication:         auth,
		NotificationApplication: notification,
	}
}

func authInjection() application.AuthApplication {
	ar := repository.NewAuthRepository()
	as := service.NewAuthService(ar)

	aa := application.NewAuthApplication(as)

	return aa
}

func notificationInjection(db *repository.Client) application.NotificationApplication {
	nr := repository.NewNotificationRepository(db)
	ndv := dv.NewNotificationDomainValidation()
	ns := service.NewNotificationService(ndv, nr)

	nrv := rv.NewNotificationRequestValidation()
	na := application.NewNotificationApplication(nrv, ns)

	return na
}
