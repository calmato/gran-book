package user

import "context"

// Service - Userドメインサービス
type Service interface {
	Create(ctx context.Context, u *User) error
}
