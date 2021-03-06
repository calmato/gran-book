package chat

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
)

// Service - Chatドメインサービス
type Service interface {
	ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*Room, error)
	CreateRoom(ctx context.Context, cr *Room) error
	ValidationRoom(ctx context.Context, cr *Room) error
	PushCreateRoom(ctx context.Context, cr *Room) error
	PushNewMessage(ctx context.Context, cr *Room, cm *Message) error
}
