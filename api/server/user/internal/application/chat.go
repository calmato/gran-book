package application

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/array"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// ChatApplication - Chatアプリケーションのインターフェース
type ChatApplication interface {
	ListRoom(ctx context.Context, userID string, q *database.ListQuery) ([]*chat.Room, error)
	GetRoom(ctx context.Context, roomID string, userID string) (*chat.Room, error)
	CreateRoom(ctx context.Context, cr *chat.Room) error
	CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error
	UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error)
}

type chatApplication struct {
	chatDomainValidation chat.Validation
	chatRepository       chat.Repository
	chatUploader         chat.Uploader
}

// NewChatApplication - ChatApplicationの生成
func NewChatApplication(cdv chat.Validation, cr chat.Repository, cu chat.Uploader) ChatApplication {
	return &chatApplication{
		chatDomainValidation: cdv,
		chatRepository:       cr,
		chatUploader:         cu,
	}
}

func (a *chatApplication) ListRoom(ctx context.Context, userID string, q *database.ListQuery) ([]*chat.Room, error) {
	q.Order = &database.OrderQuery{
		Field:   "updatedAt",
		OrderBy: database.OrderByDesc,
	}

	crs, err := a.chatRepository.ListRoom(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func (a *chatApplication) GetRoom(ctx context.Context, roomID string, userID string) (*chat.Room, error) {
	cr, err := a.chatRepository.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	isJoin, _ := array.Contains(cr.UserIDs, userID)
	if !isJoin {
		err := xerrors.New("This user is not join the room")
		return nil, exception.Forbidden.New(err)
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

	err = a.chatRepository.CreateRoom(ctx, cr)
	if err != nil {
		return err
	}

	return nil
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

	err = a.chatRepository.UpdateRoom(ctx, cr)
	if err != nil {
		return err
	}

	return nil
}

func (a *chatApplication) UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error) {
	return a.chatUploader.Image(ctx, cr.ID, image)
}
