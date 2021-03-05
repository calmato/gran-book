package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Userドメインサービス
type Service interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, int64, error)
	// ListFollow(ctx context.Context, q *domain.ListQuery) ([]*User, int64, error)
	// ListFollower(ctx context.Context, q *domain.ListQuery) ([]*User, int64, error)
	ListFriendsCount(ctx context.Context, u *User) (int64, int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	// ShowFollow(ctx context.Context, id int64) (*Follow, error)
	Create(ctx context.Context, u *User) error
	// CreateFollow(ctx context.Context, f *Follow) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error)
	// DeleteFollow(ctx context.Context, id int64) error
	IsFriend(ctx context.Context, u *User, cuid string) (bool, bool, error)
}
