package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Userドメインサービス
type Service interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollow(ctx context.Context, q *domain.ListQuery, uid string) ([]*Follow, error)
	ListFollower(ctx context.Context, q *domain.ListQuery, uid string) ([]*Follower, error)
	ListCount(ctx context.Context, q *domain.ListQuery) (int64, error)
	ListFriendCount(ctx context.Context, uid string) (int64, int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	ShowRelationship(ctx context.Context, id int64) (*Relationship, error)
	ShowRelationshipByUID(ctx context.Context, followID string, followerID string) (*Relationship, error)
	Create(ctx context.Context, u *User) error
	CreateRelationship(ctx context.Context, r *Relationship) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	DeleteRelationship(ctx context.Context, id int64) error
	UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error)
	IsFriend(ctx context.Context, friendID string, uid string) (bool, bool)
}
