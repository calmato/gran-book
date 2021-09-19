package server

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
	v2 "github.com/calmato/gran-book/api/gateway/native/internal/server/v2"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	HTTP          util.HTTPHandler
	V1Api         v1.ApiV1Handler
	V2Api         v2.ApiV2Handler
}

// Params - DIコンテナ生成用のパラメータ
type Params struct {
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

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(params *Params) (*Registry, error) {
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

	return &Registry{
		Authenticator: util.NewAuthenticator(params.FirebaseAuth),
		HTTP:          util.NewHTTPHandler(),
		V1Api:         v1.NewApiV1Handler(v1Params, datetime.Now),
		V2Api:         v2.NewApiV2Handler(v2Params, datetime.Now),
	}, nil
}

func newGRPCClient(params *Params) (*gRPCClient, error) {
	authConn, err := grpc.Dial(params.AuthServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bookConn, err := grpc.Dial(params.BookServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	chatConn, err := grpc.Dial(params.ChatServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.Dial(params.UserServiceURL, grpc.WithInsecure())
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
