package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Userドメインサービス
type Service interface {
	// TODO: refactor
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, error)
	ListFollow(ctx context.Context, q *domain.ListQuery, uid string) ([]*Follow, error)
	ListFollower(ctx context.Context, q *domain.ListQuery, uid string) ([]*Follower, error)
	ListInstanceID(ctx context.Context, userIDs []string) ([]string, error)
	ListCount(ctx context.Context, q *domain.ListQuery) (int, error)
	ListFriendCount(ctx context.Context, uid string) (int, int, error)
	Show(ctx context.Context, uid string) (*User, error)
	ShowRelationship(ctx context.Context, id int) (*Relationship, error)
	ShowRelationshipByUID(ctx context.Context, followID string, followerID string) (*Relationship, error)
	Create(ctx context.Context, u *User) error
	CreateWithOAuth(ctx context.Context, u *User) error
	CreateRelationship(ctx context.Context, r *Relationship) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	Delete(ctx context.Context, uid string) error
	DeleteRelationship(ctx context.Context, id int) error
	UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error)
	IsFriend(ctx context.Context, friendID string, uid string) (bool, bool)
	Validation(ctx context.Context, u *User) error
	ValidationRelationship(ctx context.Context, r *Relationship) error
}
