package storage

import (
	"context"

	"github.com/calmato/gran-book/api/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/pkg/firebase/storage"
)

type chatUploader struct {
	storage *storage.Storage
}

// NewChatUploader - ChatUploaderの生成
func NewChatUploader(s *storage.Storage) chat.Uploader {
	return &chatUploader{
		storage: s,
	}
}

func (u *chatUploader) Image(ctx context.Context, roomID string, data []byte) (string, error) {
	path := getChatMessagePath(roomID)

	imageURL, err := u.storage.Write(ctx, path, data)
	if err != nil {
		return "", exception.ErrInStorage.New(err)
	}

	return imageURL, nil
}
