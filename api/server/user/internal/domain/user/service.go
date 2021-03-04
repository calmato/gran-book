package user

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Userドメインサービス
type Service interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *domain.ListQuery) ([]*User, int64, error)
	ListFriendsCount(ctx context.Context, u *User) (int64, int64, error)
	Show(ctx context.Context, uid string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error)
}
