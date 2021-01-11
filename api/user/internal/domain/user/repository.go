package user

import "context"

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
