package registry

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	Health        util.HealthHandler
	V1Auth        v1.AuthHandler
	V1Chat        v1.ChatHandler
	V1User        v1.UserHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(
	fa *authentication.Auth,
	authServiceURL string, userServiceURL string, chatServiceURL string,
) (*Registry, error) {
	v1AuthConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	v1UserConn, err := grpc.Dial(userServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	v1ChatConn, err := grpc.Dial(chatServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(fa),
		Health:        util.NewHealthHandler(),
		V1Auth:        v1.NewAuthHandler(v1AuthConn),
		V1Chat:        v1.NewChatHandler(v1ChatConn, v1AuthConn, v1UserConn),
		V1User:        v1.NewUserHandler(v1UserConn),
	}, nil
}
