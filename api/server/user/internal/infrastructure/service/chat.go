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
	chatMessaging        chat.Messaging
}

// NewChatService - ChatServiceの生成
func NewChatService(cdv chat.Validation, cr chat.Repository, cm chat.Messaging) chat.Service {
	return &chatService{
		chatDomainValidation: cdv,
		chatRepository:       cr,
		chatMessaging:        cm,
	}
}

func (s *chatService) ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*chat.Room, error) {
	return s.chatRepository.ListRoom(ctx, q, uid)
}

func (s *chatService) CreateRoom(ctx context.Context, cr *chat.Room) error {
	current := time.Now().Local()

	cr.ID = uuid.New().String()
	cr.CreatedAt = current
	cr.UpdatedAt = current

	return s.chatRepository.CreateRoom(ctx, cr)
}

func (s *chatService) ValidationRoom(ctx context.Context, cr *chat.Room) error {
	return s.chatDomainValidation.Room(ctx, cr)
}

func (s *chatService) PushNewMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	return s.chatMessaging.PushNewMessage(cr, cm)
}
