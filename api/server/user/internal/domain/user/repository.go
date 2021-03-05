package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollow(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollower(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	ListFollowCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	ShowFollow(ctx context.Context, id int64) (*Follow, error)
	Create(ctx context.Context, u *User) error
	CreateFollow(ctx context.Context, f *Follow) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	DeleteFollow(ctx context.Context, f *Follow) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
	GetFollowIDByUserID(ctx context.Context, followID string, followerID string) (int64, error)
}
