package storage

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	gcs "github.com/calmato/gran-book/api/server/user/pkg/firebase/storage"
)

type chatUploader struct {
	storage *gcs.Storage
}

// NewChatUploader - ChatUploaderの生成
func NewChatUploader(s *gcs.Storage) chat.Uploader {
	return &chatUploader{
		storage: s,
	}
}

func (u *chatUploader) Image(ctx context.Context, roomID string, data []byte) (string, error) {
	path := getChatMessagePath(roomID)

	imageURL, err := u.storage.Write(ctx, path, data)
	if err != nil {
		return "", exception.ErrorInStorage.New(err)
	}

	return imageURL, nil
}
