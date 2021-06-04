package chat

import "context"

// Repository - Chatレポジトリ
type Repository interface {
	CreateRoom(ctx context.Context, cr *Room) error
}
