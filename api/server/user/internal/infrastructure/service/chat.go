package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/google/uuid"
)

type chatService struct {
	chatDomainValidation chat.Validation
	chatRepository       chat.Repository
	chatUploader         chat.Uploader
	chatMessaging        chat.Messaging
}

// NewChatService - ChatServiceの生成
func NewChatService(cdv chat.Validation, cr chat.Repository, cu chat.Uploader, cm chat.Messaging) chat.Service {
	return &chatService{
		chatDomainValidation: cdv,
		chatRepository:       cr,
		chatUploader:         cu,
		chatMessaging:        cm,
	}
}

func (s *chatService) ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*chat.Room, error) {
	return s.chatRepository.ListRoom(ctx, q, uid)
}

func (s *chatService) GetRoom(ctx context.Context, roomID string) (*chat.Room, error) {
	return s.chatRepository.GetRoom(ctx, roomID)
}

func (s *chatService) CreateRoom(ctx context.Context, cr *chat.Room) error {
	current := time.Now().Local()

	cr.ID = uuid.New().String()
	cr.CreatedAt = current
	cr.UpdatedAt = current

	return s.chatRepository.CreateRoom(ctx, cr)
}

func (s *chatService) CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	current := time.Now().Local()

	cm.ID = uuid.New().String()
	cm.CreatedAt = current
	cr.LatestMessage = cm
	cr.UpdatedAt = current

	err := s.chatRepository.CreateMessage(ctx, cr.ID, cm)
	if err != nil {
		return err
	}

	return s.chatRepository.UpdateRoom(ctx, cr)
}

func (s *chatService) ValidationRoom(ctx context.Context, cr *chat.Room) error {
	return s.chatDomainValidation.Room(ctx, cr)
}

func (s *chatService) ValidationMessage(ctx context.Context, cm *chat.Message) error {
	return s.chatDomainValidation.Message(ctx, cm)
}

func (s *chatService) UploadImage(ctx context.Context, roomID string, image []byte) (string, error) {
	return s.chatUploader.Image(ctx, roomID, image)
}

func (s *chatService) PushCreateRoom(ctx context.Context, cr *chat.Room) error {
	return s.chatMessaging.PushCreateRoom(cr)
}

func (s *chatService) PushNewMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	return s.chatMessaging.PushNewMessage(cr, cm)
}
