package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollows(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollowers(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	ListFollowsCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	ListFollowersCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
