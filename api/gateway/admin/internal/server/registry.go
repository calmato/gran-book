package server

import (
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
)

// Registry - DIコンテナ
type Registry struct {
	Authenticator util.Authenticator
	Health        util.HealthHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry(fa *authentication.Auth) (*Registry, error) {
	return &Registry{
		Authenticator: util.NewAuthenticator(fa),
		Health:        util.NewHealthHandler(),
	}, nil
}
