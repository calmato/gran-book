package auth

import "context"

// Repository - Authレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
}
