package server

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
	v2 "github.com/calmato/gran-book/api/gateway/native/internal/server/v2"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/gateway/native/proto/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/user"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	HTTP          util.HTTPHandler
	V1Auth        v1.AuthHandler
	V1Book        v1.BookHandler
	V1Bookshelf   v1.BookshelfHandler
	V1Chat        v1.ChatHandler
	V1Review      v1.ReviewHandler
	V1Top         v1.TopHandler
	V1User        v1.UserHandler
	V2Book        v2.BookHandler
	V2Bookshelf   v2.BookshelfHandler
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
	client, err := newgRPCClient(params)
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(params.FirebaseAuth),
		HTTP:          util.NewHTTPHandler(),
		V1Auth:        v1.NewAuthHandler(client.auth),
		V1Book:        v1.NewBookHandler(client.auth, client.book),
		V1Bookshelf:   v1.NewBookshelfHandler(client.auth, client.book),
		V1Chat:        v1.NewChatHandler(client.auth, client.chat, client.user),
		V1Review:      v1.NewReviewHandler(client.auth, client.book, client.user),
		V1Top:         v1.NewTopHandler(client.auth, client.book),
		V1User:        v1.NewUserHandler(client.user),
		V2Book:        v2.NewBookHandler(client.auth, client.book, client.user),
		V2Bookshelf:   v2.NewBookshelfHandler(client.book, client.user),
	}, nil
}

func newgRPCClient(params *Params) (*gRPCClient, error) {
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
