package chat

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Chatドメインサービス
type Service interface {
	ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*Room, error)
	GetRoom(ctx context.Context, roomID string) (*Room, error)
	CreateRoom(ctx context.Context, cr *Room) error
	CreateMessage(ctx context.Context, cr *Room, cm *Message) error
	ValidationRoom(ctx context.Context, cr *Room) error
	UploadImage(ctx context.Context, roomID string, image []byte) (string, error)
	PushCreateRoom(ctx context.Context, cr *Room) error
	PushNewMessage(ctx context.Context, cr *Room, cm *Message) error
}
