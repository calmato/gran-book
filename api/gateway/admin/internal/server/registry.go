package server

import (
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/admin/internal/server/v1"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/user"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	HTTP          util.HTTPHandler
	V1Admin       v1.AdminHandler
	V1Auth        v1.AuthHandler
	V1Book        v1.BookHandler
	V1User        v1.UserHandler
}

// Params - DIコンテナ生成用のパラメータ
type Params struct {
	FirebaseAuth    *authentication.Auth
	AdminServiceURL string
	AuthServiceURL  string
	BookServiceURL  string
	UserServiceURL  string
}

type gRPCClient struct {
	admin user.AdminServiceClient
	auth  user.AuthServiceClient
	book  book.BookServiceClient
	user  user.UserServiceClient
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(params *Params) (*Registry, error) {
	client, err := newgRPCClient(params)
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(params.FirebaseAuth, client.auth),
		HTTP:          util.NewHTTPHandler(),
		V1Admin:       v1.NewAdminHandler(client.admin),
		V1Auth:        v1.NewAuthHandler(client.auth),
		V1Book:        v1.NewBookHandler(client.book),
		V1User:        v1.NewUserHandler(client.user),
	}, nil
}

func newgRPCClient(params *Params) (*gRPCClient, error) {
	adminConn, err := grpc.Dial(params.AdminServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authConn, err := grpc.Dial(params.AuthServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bookConn, err := grpc.Dial(params.BookServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.Dial(params.UserServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &gRPCClient{
		admin: user.NewAdminServiceClient(adminConn),
		auth:  user.NewAuthServiceClient(authConn),
		user:  user.NewUserServiceClient(userConn),
		book:  book.NewBookServiceClient(bookConn),
	}, nil
}
