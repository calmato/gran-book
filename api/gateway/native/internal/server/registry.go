package server

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
	v2 "github.com/calmato/gran-book/api/gateway/native/internal/server/v2"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	Health        util.HealthHandler
	V1Auth        v1.AuthHandler
	V1Book        v1.BookHandler
	V1Bookshelf   v1.BookshelfHandler
	V1Chat        v1.ChatHandler
	V1Review      v1.ReviewHandler
	V1User        v1.UserHandler
	V2Book        v2.BookHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(
	fa *authentication.Auth,
	authServiceURL string, userServiceURL string, chatServiceURL string, bookServiceURL string,
) (*Registry, error) {
	authConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.Dial(userServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	chatConn, err := grpc.Dial(chatServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	bookConn, err := grpc.Dial(bookServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(fa),
		Health:        util.NewHealthHandler(),
		V1Auth:        v1.NewAuthHandler(authConn),
		V1Book:        v1.NewBookHandler(bookConn, authConn),
		V1Bookshelf:   v1.NewBookshelfHandler(bookConn, authConn, userConn),
		V1Chat:        v1.NewChatHandler(chatConn, authConn, userConn),
		V1Review:      v1.NewReviewHandler(bookConn, authConn, userConn),
		V1User:        v1.NewUserHandler(userConn),
		V2Book:        v2.NewBookHandler(bookConn, authConn, userConn),
	}, nil
}
