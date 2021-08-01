package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/application"
	rv "github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/messaging"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/service"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/server/user/internal/infrastructure/validation"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/server/user/pkg/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AdminRequestValidation validation.AdminRequestValidation
	AuthRequsetValidation  validation.AuthRequestValidation
	ChatApplication        application.ChatApplication
	UserApplication        application.UserApplication
	UserRequestValidation  validation.UserRequestValidation
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(
	db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) *Registry {
	adminRequestValidation := adminInjection()
	authRequestValidation := authInjection()
	chatApplication := chatInjection(db, fa, fs, s)
	userApplication, userRequestValidation := userInjection(db, fa, s)

	return &Registry{
		AdminRequestValidation: adminRequestValidation,
		AuthRequsetValidation:  authRequestValidation,
		ChatApplication:        chatApplication,
		UserApplication:        userApplication,
		UserRequestValidation:  userRequestValidation,
	}
}

func adminInjection() validation.AdminRequestValidation {
	arv := validation.NewAdminRequestValidation()
	return arv
}

func authInjection() validation.AuthRequestValidation {
	arv := validation.NewAuthRequestValidation()
	return arv
}

func chatInjection(
	db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) application.ChatApplication {
	cm := messaging.NewChatMessaging()
	cr := repository.NewChatRepository(fs)
	cdv := dv.NewChatDomainValidation()
	cu := storage.NewChatUploader(s)
	cs := service.NewChatService(cdv, cr, cu, cm)

	crv := rv.NewChatRequestValidation()
	ca := application.NewChatApplication(crv, cs)

	return ca
}

func userInjection(
	db *database.Client, fa *authentication.Auth, s *gcs.Storage,
) (application.UserApplication, validation.UserRequestValidation) {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	ua := application.NewUserApplication(udv, ur, uu)

	urv := validation.NewUserRequestValidation()

	return ua, urv
}
