package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, query *domain.ListQuery) ([]*User, error)
	Show(ctx context.Context, uid string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
