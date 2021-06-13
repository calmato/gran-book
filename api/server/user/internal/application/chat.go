package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
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
}

// NewChatApplication - ChatApplicationの生成
func NewChatApplication(crv validation.ChatRequestValidation, cs chat.Service) ChatApplication {
	return &chatApplication{
		chatRequestValidation: crv,
		chatService:           cs,
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

	err = a.chatService.CreateRoom(ctx, cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}
