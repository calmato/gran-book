package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/domain/auth"
)

// AuthApplication - Authアプリケーションのインターフェース
type AuthApplication interface {
	Authentication(ctx context.Context) (string, error)
}

type authApplication struct {
	authService auth.Service
}

// NewAuthApplication - AuthApplicationの生成
func NewAuthApplication(as auth.Service) AuthApplication {
	return &authApplication{
		authService: as,
	}
}

func (a *authApplication) Authentication(ctx context.Context) (string, error) {
	return a.authService.Authentication(ctx)
}
