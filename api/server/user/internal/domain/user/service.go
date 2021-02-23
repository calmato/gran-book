package user

import "context"

// Service - Userドメインサービス
type Service interface {
	Authentication(ctx context.Context) (string, error)
	Show(ctx context.Context, uid string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, uid string, password string) error
	UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error)
}
