package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/application"
	rv "github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/messaging"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/server/user/internal/infrastructure/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/server/user/pkg/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AdminApplication application.AdminApplication
	UserApplication  application.UserApplication
	ChatApplication  application.ChatApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(
	db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) *Registry {
	admin := adminInjection(db, fa, s)
	user := userInjection(db, fa, s)
	chat := chatInjection(db, fa, fs, s)

	return &Registry{
		AdminApplication: admin,
		UserApplication:  user,
		ChatApplication:  chat,
	}
}

func adminInjection(db *database.Client, fa *authentication.Auth, s *gcs.Storage) application.AdminApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	arv := rv.NewAdminRequestValidation()
	aa := application.NewAdminApplication(arv, us)

	return aa
}

func userInjection(db *database.Client, fa *authentication.Auth, s *gcs.Storage) application.UserApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)

	ua := application.NewUserApplication(udv, ur, uu)
	return ua
}

func chatInjection(
	db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) application.ChatApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	cm := messaging.NewChatMessaging()
	cr := repository.NewChatRepository(fs)
	cdv := dv.NewChatDomainValidation()
	cu := storage.NewChatUploader(s)
	cs := service.NewChatService(cdv, cr, cu, cm)

	crv := rv.NewChatRequestValidation()
	ca := application.NewChatApplication(crv, cs, us)

	return ca
}
