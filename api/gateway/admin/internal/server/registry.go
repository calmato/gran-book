package server

import (
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/admin/internal/server/v1"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	Health        util.HealthHandler
	V1Admin       v1.AdminHandler
	V1Auth        v1.AuthHandler
	V1Book        v1.BookHandler
	V1User        v1.UserHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(
	fa *authentication.Auth,
	authServiceURL string, adminServiceURL string, userServiceURL string, bookServiceURL string,
) (*Registry, error) {
	authConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	adminConn, err := grpc.Dial(adminServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.Dial(userServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	bookConn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(fa, authConn),
		Health:        util.NewHealthHandler(),
		V1Admin:       v1.NewAdminHandler(adminConn),
		V1Auth:        v1.NewAuthHandler(authConn),
		V1Book:        v1.NewBookHandler(bookConn),
		V1User:        v1.NewUserHandler(userConn),
	}, nil
}
