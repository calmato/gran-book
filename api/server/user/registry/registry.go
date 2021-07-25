package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/application"
	rv "github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/messaging"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/server/user/internal/infrastructure/validation"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/server/user/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AdminApplication application.AdminApplication
	AuthApplication  application.AuthApplication
	UserApplication  application.UserApplication
	ChatApplication  application.ChatApplication
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(
	db *repository.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) *Registry {
	admin := adminInjection(db, fa, s)
	auth := authInjection(db, fa, s)
	user := userInjection(db, fa, s)
	chat := chatInjection(db, fa, fs, s)

	return &Registry{
		AdminApplication: admin,
		AuthApplication:  auth,
		UserApplication:  user,
		ChatApplication:  chat,
	}
}

func adminInjection(db *repository.Client, fa *authentication.Auth, s *gcs.Storage) application.AdminApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	arv := rv.NewAdminRequestValidation()
	aa := application.NewAdminApplication(arv, us)

	return aa
}

func authInjection(db *repository.Client, fa *authentication.Auth, s *gcs.Storage) application.AuthApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	arv := rv.NewAuthRequestValidation()
	aa := application.NewAuthApplication(arv, us)

	return aa
}

func userInjection(db *repository.Client, fa *authentication.Auth, s *gcs.Storage) application.UserApplication {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	us := service.NewUserService(udv, ur, uu)

	urv := rv.NewUserRequestValidation()
	ua := application.NewUserApplication(urv, us)

	return ua
}

func chatInjection(
	db *repository.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
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
