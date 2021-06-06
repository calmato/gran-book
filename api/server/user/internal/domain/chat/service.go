package chat

import "context"

// Service - Chatドメインサービス
type Service interface {
	CreateRoom(ctx context.Context, cr *Room) error
	ValidationRoom(ctx context.Context, cr *Room) error
}
