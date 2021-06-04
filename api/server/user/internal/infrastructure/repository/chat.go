package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/firestore"
)

type chatRepository struct {
	firestore *firestore.Firestore
}

// NewChatRepository - ChatRepositoryの生成
func NewChatRepository(fs *firestore.Firestore) chat.Repository {
	return &chatRepository{
		firestore: fs,
	}
}

func (r *chatRepository) CreateRoom(ctx context.Context, cr *chat.Room) error {
	c := getChatRoomCollection()

	err := r.firestore.Set(ctx, c, cr.ID, cr)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
