package native

import (
	"crypto/tls"
	"crypto/x509"

	v1 "github.com/calmato/gran-book/api/internal/gateway/native/v1/handler"
	v2 "github.com/calmato/gran-book/api/internal/gateway/native/v2/handler"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/calmato/gran-book/api/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type registry struct {
	Authenticator util.Authenticator
	V1Api         v1.APIV1Handler
	V2Api         v2.APIV2Handler
}

// params - DIコンテナ生成用のパラメータ
type params struct {
	Insecure       bool
	FirebaseAuth   *authentication.Auth
	AuthServiceURL string
	BookServiceURL string
	ChatServiceURL string
	UserServiceURL string
}

type gRPCClient struct {
	auth user.AuthServiceClient
	book book.BookServiceClient
	chat chat.ChatServiceClient
	user user.UserServiceClient
}

func newRegistry(params *params) (*registry, error) {
	client, err := newGRPCClient(params)
	if err != nil {
		return nil, err
	}

	v1Params := &v1.Params{
		Auth: client.auth,
		User: client.user,
		Chat: client.chat,
		Book: client.book,
	}
	v2Params := &v2.Params{
		Auth: client.auth,
		User: client.user,
		Book: client.book,
	}

	return &registry{
		Authenticator: util.NewAuthenticator(params.FirebaseAuth),
		V1Api:         v1.NewAPIV1Handler(v1Params, datetime.Now),
		V2Api:         v2.NewAPIV2Handler(v2Params, datetime.Now),
	}, nil
}

func newGRPCClient(params *params) (*gRPCClient, error) {
	opts, err := newGRPCOptios(params)
	if err != nil {
		return nil, err
	}

	authConn, err := grpc.Dial(params.AuthServiceURL, opts)
	if err != nil {
		return nil, err
	}
	bookConn, err := grpc.Dial(params.BookServiceURL, opts)
	if err != nil {
		return nil, err
	}
	chatConn, err := grpc.Dial(params.ChatServiceURL, opts)
	if err != nil {
		return nil, err
	}
	userConn, err := grpc.Dial(params.UserServiceURL, opts)
	if err != nil {
		return nil, err
	}

	return &gRPCClient{
		auth: user.NewAuthServiceClient(authConn),
		user: user.NewUserServiceClient(userConn),
		chat: chat.NewChatServiceClient(chatConn),
		book: book.NewBookServiceClient(bookConn),
	}, nil
}

func newGRPCOptios(params *params) ([]grpc.DialOption, error) {
	var opts []grpc.DialOption
	if params.Insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return opts, nil
}
