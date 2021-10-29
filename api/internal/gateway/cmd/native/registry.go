package native

import (
	v1 "github.com/calmato/gran-book/api/internal/gateway/native/v1/handler"
	v2 "github.com/calmato/gran-book/api/internal/gateway/native/v2/handler"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/calmato/gran-book/api/proto/user"
	"google.golang.org/grpc"
)

type registry struct {
	Authenticator util.Authenticator
	V1Api         v1.APIV1Handler
	V2Api         v2.APIV2Handler
}

// params - DIコンテナ生成用のパラメータ
type params struct {
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
