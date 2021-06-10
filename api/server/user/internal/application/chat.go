package application

import (
	"context"
	"log"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/array"
)

// ChatApplication - Chatアプリケーションのインターフェース
type ChatApplication interface {
	ListRoom(ctx context.Context, cuid string) ([]*chat.Room, error)
	CreateRoom(ctx context.Context, in *input.CreateRoom, cuid string) (*chat.Room, error)
}

type chatApplication struct {
	chatRequestValidation validation.ChatRequestValidation
	chatService           chat.Service
	userService           user.Service
}

// NewChatApplication - ChatApplicationの生成
func NewChatApplication(crv validation.ChatRequestValidation, cs chat.Service, us user.Service) ChatApplication {
	return &chatApplication{
		chatRequestValidation: crv,
		chatService:           cs,
		userService:           us,
	}
}

func (a *chatApplication) ListRoom(ctx context.Context, cuid string) ([]*chat.Room, error) {
	q := &domain.ListQuery{
		Order: &domain.QueryOrder{
			By:        "updatedAt",
			Direction: "desc",
		},
	}

	return a.chatService.ListRoom(ctx, q, cuid)
}

func (a *chatApplication) CreateRoom(ctx context.Context, in *input.CreateRoom, cuid string) (*chat.Room, error) {
	err := a.chatRequestValidation.CreateRoom(in)
	if err != nil {
		return nil, err
	}

	cr := &chat.Room{
		UserIDs: in.UserIDs,
	}

	if ok, _ := array.Contains(cr.UserIDs, cuid); !ok {
		cr.UserIDs = append(cr.UserIDs, cuid)
	}

	err = a.chatService.ValidationRoom(ctx, cr)
	if err != nil {
		return nil, err
	}

	instanceIDs, err := a.userService.ListInstanceID(ctx, cr.UserIDs)
	if err != nil {
		return nil, err
	}
	cr.InstanceIDs = instanceIDs

	err = a.chatService.CreateRoom(ctx, cr)
	if err != nil {
		return nil, err
	}

	err = a.chatService.PushCreateRoom(ctx, cr)
	if err != nil {
		log.Printf("Failed to push notification: %v", err) // TOOD: エラーの出し方考える
	}

	return cr, nil
}
