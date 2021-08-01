package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/pkg/array"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
)

// ChatApplication - Chatアプリケーションのインターフェース
type ChatApplication interface {
	ListRoom(ctx context.Context, cuid string) ([]*chat.Room, error)
	CreateRoom(ctx context.Context, in *input.CreateRoom, cuid string) (*chat.Room, error)
	CreateTextMessage(ctx context.Context, in *input.CreateTextMessage, roomID string, cuid string) (*chat.Message, error)
	CreateImageMessage(
		ctx context.Context, in *input.CreateImageMessage, roomID string, cuid string,
	) (*chat.Message, error)
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
	q := &database.ListQuery{
		Order: &database.OrderQuery{
			Field:   "updatedAt",
			OrderBy: database.OrderByDesc,
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

func (a *chatApplication) CreateTextMessage(
	ctx context.Context, in *input.CreateTextMessage, roomID string, cuid string,
) (*chat.Message, error) {
	err := a.chatRequestValidation.CreateTextMessage(in)
	if err != nil {
		return nil, err
	}

	cr, err := a.chatService.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	cm := &chat.Message{
		Text:   in.Text,
		UserID: cuid,
	}

	err = a.chatService.ValidationMessage(ctx, cm)
	if err != nil {
		return nil, err
	}

	err = a.chatService.CreateMessage(ctx, cr, cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (a *chatApplication) CreateImageMessage(
	ctx context.Context, in *input.CreateImageMessage, roomID string, cuid string,
) (*chat.Message, error) {
	err := a.chatRequestValidation.CreateImageMessage(in)
	if err != nil {
		return nil, err
	}

	cr, err := a.chatService.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	imageURL, err := a.chatService.UploadImage(ctx, cr.ID, in.Image)
	if err != nil {
		return nil, err
	}

	cm := &chat.Message{
		Image:  imageURL,
		UserID: cuid,
	}

	err = a.chatService.ValidationMessage(ctx, cm)
	if err != nil {
		return nil, err
	}

	err = a.chatService.CreateMessage(ctx, cr, cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}
