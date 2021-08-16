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
	V1Auth        v1.AuthHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(fa *authentication.Auth, authServiceURL string) (*Registry, error) {
	authConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Registry{
		Authenticator: util.NewAuthenticator(fa),
		Health:        util.NewHealthHandler(),
		V1Auth:        v1.NewAuthHandler(authConn),
	}, nil
}
