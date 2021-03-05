package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollow(ctx context.Context, q *domain.ListQuery) ([]*Follow, error)
	ListFollower(ctx context.Context, q *domain.ListQuery) ([]*Follower, error)
	ListCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	ListRelationshipCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	ShowRelationship(ctx context.Context, id int64) (*Relationship, error)
	Create(ctx context.Context, u *User) error
	CreateRelationship(ctx context.Context, r *Relationship) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	DeleteRelationship(ctx context.Context, id int64) error
	GetUIDByEmail(ctx context.Context, email string) (string, error)
}
