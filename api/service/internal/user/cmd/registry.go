package cmd

import (
	"github.com/calmato/gran-book/api/service/internal/user/application"
	"github.com/calmato/gran-book/api/service/internal/user/infrastructure/repository"
	"github.com/calmato/gran-book/api/service/internal/user/infrastructure/storage"
	dv "github.com/calmato/gran-book/api/service/internal/user/infrastructure/validation"
	"github.com/calmato/gran-book/api/service/internal/user/interface/server"
	rv "github.com/calmato/gran-book/api/service/internal/user/interface/validation"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"github.com/calmato/gran-book/api/service/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/service/pkg/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/service/pkg/firebase/storage"
	"github.com/calmato/gran-book/api/service/proto/chat"
	"github.com/calmato/gran-book/api/service/proto/user"
)

type registry struct {
	admin user.AdminServiceServer
	auth  user.AuthServiceServer
	chat  chat.ChatServiceServer
	user  user.UserServiceServer
}

func newRegistry(db *database.Client, fa *authentication.Auth, fs *firestore.Firestore, s *gcs.Storage) *registry {
	adminRequestValidation := adminInjector()
	authRequestValidation := authInjector()
	chatApplication, chatRequestValidation := chatInjector(fs, s)
	userApplication, userRequestValidation := userInjector(db, fa, s)

	return &registry{
		admin: server.NewAdminServer(adminRequestValidation, userApplication),
		auth:  server.NewAuthServer(authRequestValidation, userApplication),
		chat:  server.NewChatServer(chatRequestValidation, chatApplication),
		user:  server.NewUserServer(userRequestValidation, userApplication),
	}
}

func adminInjector() rv.AdminRequestValidation {
	arv := rv.NewAdminRequestValidation()
	return arv
}

func authInjector() rv.AuthRequestValidation {
	arv := rv.NewAuthRequestValidation()
	return arv
}

func chatInjector(
	fs *firestore.Firestore, s *gcs.Storage,
) (application.ChatApplication, rv.ChatRequestValidation) {
	cr := repository.NewChatRepository(fs)
	cdv := dv.NewChatDomainValidation()
	cu := storage.NewChatUploader(s)
	ca := application.NewChatApplication(cdv, cr, cu)

	crv := rv.NewChatRequestValidation()

	return ca, crv
}

func userInjector(
	db *database.Client, fa *authentication.Auth, s *gcs.Storage,
) (application.UserApplication, rv.UserRequestValidation) {
	ur := repository.NewUserRepository(db, fa)
	udv := dv.NewUserDomainValidation(ur)
	uu := storage.NewUserUploader(s)
	ua := application.NewUserApplication(udv, ur, uu)

	urv := rv.NewUserRequestValidation()

	return ua, urv
}
