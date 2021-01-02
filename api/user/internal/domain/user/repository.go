package user

import "context"

// Repository - Userレポジトリ
type Repository interface {
	Create(ctx context.Context, u *User) error
}
