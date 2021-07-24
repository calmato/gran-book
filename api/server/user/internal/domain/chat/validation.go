package chat

import "context"

// Validation - Chatドメインバリデーション
type Validation interface {
	Room(ctx context.Context, cr *Room) error
	Message(ctx context.Context, cm *Message) error
}
