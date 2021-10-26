package admin

import (
	v1 "github.com/calmato/gran-book/api/internal/gateway/admin/handler/v1"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/user"
	"google.golang.org/grpc"
)

type registry struct {
	Authenticator util.Authenticator
	V1Api         v1.APIV1Handler
}

// params - DIコンテナ生成用のパラメータ
type params struct {
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

func newRegistry(params *params) (*registry, error) {
	client, err := newGRPCClient(params)
	if err != nil {
		return nil, err
	}

	v1Params := &v1.Params{
		Auth:  client.auth,
		User:  client.user,
		Admin: client.admin,
		Book:  client.book,
	}

	return &registry{
		Authenticator: util.NewAuthenticator(params.FirebaseAuth),
		V1Api:         v1.NewAPIV1Handler(v1Params, datetime.Now),
	}, nil
}

func newGRPCClient(params *params) (*gRPCClient, error) {
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
