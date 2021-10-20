package application

import (
	"context"
	"errors"
	"time"

	"github.com/calmato/gran-book/api/service/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/service/pkg/array"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/pkg/firebase/firestore"
	"github.com/google/uuid"
)

type chatApplication struct {
	chatDomainValidation chat.Validation
	chatRepository       chat.Repository
	chatUploader         chat.Uploader
}

var errNotJoinUserInRoom = errors.New("application: this user is not join the room")

// NewChatApplication - ChatApplicationの生成
func NewChatApplication(cdv chat.Validation, cr chat.Repository, cu chat.Uploader) ChatApplication {
	return &chatApplication{
		chatDomainValidation: cdv,
		chatRepository:       cr,
		chatUploader:         cu,
	}
}

func (a *chatApplication) ListRoom(ctx context.Context, userID string, p *firestore.Params) (chat.Rooms, error) {
	qs := []*firestore.Query{
		{
			Field:    "users",
			Operator: "array-contains",
			Value:    userID,
		},
	}

	return a.chatRepository.ListRoom(ctx, p, qs)
}

func (a *chatApplication) GetRoom(ctx context.Context, roomID string, userID string) (*chat.Room, error) {
	cr, err := a.chatRepository.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	isJoin, _ := array.Contains(cr.UserIDs, userID)
	if !isJoin {
		return nil, exception.ErrForbidden.New(errNotJoinUserInRoom)
	}

	return cr, nil
}

func (a *chatApplication) CreateRoom(ctx context.Context, cr *chat.Room) error {
	err := a.chatDomainValidation.Room(ctx, cr)
	if err != nil {
		return err
	}

	current := time.Now().Local()
	cr.CreatedAt = current
	cr.UpdatedAt = current
	cr.ID = uuid.New().String()

	return a.chatRepository.CreateRoom(ctx, cr)
}

func (a *chatApplication) CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	err := a.chatDomainValidation.Message(ctx, cm)
	if err != nil {
		return err
	}

	cm.CreatedAt = time.Now().Local()
	cm.ID = uuid.New().String()

	err = a.chatRepository.CreateMessage(ctx, cr.ID, cm)
	if err != nil {
		return err
	}

	cr.LatestMessage = cm
	cr.UpdatedAt = cm.CreatedAt

	return a.chatRepository.UpdateRoom(ctx, cr)
}

func (a *chatApplication) UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error) {
	return a.chatUploader.Image(ctx, cr.ID, image)
}
