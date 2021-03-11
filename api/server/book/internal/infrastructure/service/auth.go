package service

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain/auth"
)

type authService struct {
	authRepository auth.Repository
}

// NewAuthService - AuthServiceの生成
func NewAuthService(ar auth.Repository) auth.Service {
	return &authService{
		authRepository: ar,
	}
}

func (s *authService) Authentication(ctx context.Context) (string, error) {
	return s.authRepository.Authentication(ctx)
}
