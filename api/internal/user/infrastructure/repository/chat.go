package repository

import (
	"context"

	"github.com/calmato/gran-book/api/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/pkg/firebase/firestore"
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

func (r *chatRepository) ListRoom(
	ctx context.Context, p *firestore.Params, qs []*firestore.Query,
) (chat.Rooms, error) {
	c := getChatRoomCollection()

	docs, err := r.firestore.List(ctx, c, p, qs)
	if err != nil {
		return nil, exception.ToFirebaseError(err)
	}

	crs := make(chat.Rooms, len(docs))
	for i, doc := range docs {
		cr := &chat.Room{}

		err = doc.DataTo(cr)
		if err != nil {
			return nil, exception.ToFirebaseError(err)
		}

		crs[i] = cr
	}

	return crs, nil
}

func (r *chatRepository) GetRoom(ctx context.Context, roomID string) (*chat.Room, error) {
	c := getChatRoomCollection()
	cr := &chat.Room{}

	doc, err := r.firestore.Get(ctx, c, roomID)
	if err != nil {
		return nil, exception.ToFirebaseError(err)
	}

	err = doc.DataTo(cr)
	return cr, exception.ToFirebaseError(err)
}

func (r *chatRepository) CreateRoom(ctx context.Context, cr *chat.Room) error {
	c := getChatRoomCollection()

	err := r.firestore.Set(ctx, c, cr.ID, cr)
	return exception.ToFirebaseError(err)
}

func (r *chatRepository) UpdateRoom(ctx context.Context, cr *chat.Room) error {
	c := getChatRoomCollection()

	err := r.firestore.Set(ctx, c, cr.ID, cr)
	return exception.ToFirebaseError(err)
}

func (r *chatRepository) CreateMessage(ctx context.Context, roomID string, cm *chat.Message) error {
	c := getChatMessageCollection(roomID)

	err := r.firestore.Set(ctx, c, cm.ID, cm)
	return exception.ToFirebaseError(err)
}
