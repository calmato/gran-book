package registry

import (
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
	"google.golang.org/grpc"
)

// Registry - DIコンテナ
type Registry struct {
	V1Auth v1.AuthHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(authServiceURL string) (*Registry, error) {
	v1AuthConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Registry{
		V1Auth: v1.NewAuthHandler(v1AuthConn),
	}, nil
}
