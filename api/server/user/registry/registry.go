package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/server/user/internal/infrastructure/validation"
	rv "github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/server/user/pkg/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct {
	AdminRequestValidation rv.AdminRequestValidation
	AuthRequsetValidation  rv.AuthRequestValidation
	ChatApplication        application.ChatApplication
	ChatRequestValidation  rv.ChatRequestValidation
	UserApplication        application.UserApplication
	UserRequestValidation  rv.UserRequestValidation
}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(
	db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage,
) *Registry {
	adminRequestValidation := adminInjection()
	authRequestValidation := authInjection()
	chatApplication, chatRequestValidation := chatInjection(fs, s)
	userApplication, userRequestValidation := userInjection(db, fa, s)

	return &Registry{
		AdminRequestValidation: adminRequestValidation,
		AuthRequsetValidation:  authRequestValidation,
		ChatApplication:        chatApplication,
		ChatRequestValidation:  chatRequestValidation,
		UserApplication:        userApplication,
		UserRequestValidation:  userRequestValidation,
	}
}

func adminInjection() rv.AdminRequestValidation {
	arv := rv.NewAdminRequestValidation()
	return arv
}

func authInjection() rv.AuthRequestValidation {
	arv := rv.NewAuthRequestValidation()
	return arv
}

func chatInjection(
	fs *firestore.Firestore, s *gcs.Storage,
) (application.ChatApplication, rv.ChatRequestValidation) {
	cr := repository.NewChatRepository(fs)
	cdv := dv.NewChatDomainValidation()
	cu := storage.NewChatUploader(s)
	ca := application.NewChatApplication(cdv, cr, cu)

	crv := rv.NewChatRequestValidation()

	return ca, crv
}

func userInjection(
	db *database.Client, fa *authentication.Auth, s *gcs.Storage,
) (application.UserApplication, rv.UserRequestValidation) {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	ua := application.NewUserApplication(udv, ur, uu)

	urv := rv.NewUserRequestValidation()

	return ua, urv
}
