package chat

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/pkg/database"
)

// Repository - Chatレポジトリ
type Repository interface {
	ListRoom(ctx context.Context, q *database.ListQuery, uid string) ([]*Room, error)
	GetRoom(ctx context.Context, roomID string) (*Room, error)
	CreateRoom(ctx context.Context, cr *Room) error
	UpdateRoom(ctx context.Context, cr *Room) error
	CreateMessage(ctx context.Context, roomID string, cm *Message) error
}
