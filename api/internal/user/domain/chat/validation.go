//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package chat

import "context"

// Validation - Chatドメインバリデーション
type Validation interface {
	Room(ctx context.Context, cr *Room) error
	Message(ctx context.Context, cm *Message) error
}
