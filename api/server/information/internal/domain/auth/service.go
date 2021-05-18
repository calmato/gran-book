package auth

import "context"

// Service - Authドメインサービス
type Service interface {
	Authentication(ctx context.Context) (string, error)
}
