package registry

import (
	v1 "github.com/calmato/gran-book/api/gateway/native/internal/server/v1"
)

// Registry - DIコンテナ
type Registry struct {
	V1Auth v1.AuthHandler
}

// NewRegistry - internalディレクトリ配下の依存関係の解決
func NewRegistry() *Registry {
	return &Registry{
		V1Auth: v1.NewAuthHandler(),
	}
}
