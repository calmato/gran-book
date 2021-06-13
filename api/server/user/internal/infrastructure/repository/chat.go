package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
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

func (r *chatRepository) ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*chat.Room, error) {
	c := getChatRoomCollection()

	params := &firestore.Params{
		OrderBy: q.Order.By,
		SortBy:  q.Order.Direction,
	}

	qs := []*firestore.Query{
		{
			Field:    "users",
			Operator: "array-contains",
			Value:    uid,
		},
	}

	docs, err := r.firestore.List(ctx, c, params, qs)
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	crs := make([]*chat.Room, len(docs))
	for i, doc := range docs {
		cr := &chat.Room{}

		err = doc.DataTo(cr)
		if err != nil {
			return nil, err
		}

		crs[i] = cr
	}

	return crs, nil
}

func (r *chatRepository) CreateRoom(ctx context.Context, cr *chat.Room) error {
	c := getChatRoomCollection()

	err := r.firestore.Set(ctx, c, cr.ID, cr)
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
