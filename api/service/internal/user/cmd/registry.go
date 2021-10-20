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
	"github.com/calmato/gran-book/api/service/proto/user"
)

type registry struct {
	admin user.AdminServiceServer
	auth  user.AuthServiceServer
	user  user.UserServiceServer
}

func newRegistry(db *database.Client, fa *authentication.Auth, _ *firestore.Firestore, s *gcs.Storage) *registry {
	adminRequestValidation := adminInjector()
	authRequestValidation := authInjector()
	userApplication, userRequestValidation := userInjector(db, fa, s)

	return &registry{
		admin: server.NewAdminServer(adminRequestValidation, userApplication),
		auth:  server.NewAuthServer(authRequestValidation, userApplication),
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
