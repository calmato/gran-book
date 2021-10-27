//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package chat

import (
	"context"

	"github.com/calmato/gran-book/api/pkg/firebase/firestore"
)

// Repository - Chatレポジトリ
type Repository interface {
	ListRoom(ctx context.Context, p *firestore.Params, qs []*firestore.Query) (Rooms, error)
	GetRoom(ctx context.Context, roomID string) (*Room, error)
	CreateRoom(ctx context.Context, cr *Room) error
	UpdateRoom(ctx context.Context, cr *Room) error
	CreateMessage(ctx context.Context, roomID string, cm *Message) error
}
