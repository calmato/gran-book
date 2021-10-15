package storage

import (
	"context"

	"github.com/calmato/gran-book/api/service/internal/user/domain/user"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/pkg/firebase/storage"
)

type userUploader struct {
	storage *storage.Storage
}

// NewUserUploader - UserUploaderの生成
func NewUserUploader(s *storage.Storage) user.Uploader {
	return &userUploader{
		storage: s,
	}
}

func (s *userUploader) Thumbnail(ctx context.Context, userID string, data []byte) (string, error) {
	path := getUserThumbnailPath(userID)

	thumbnailURL, err := s.storage.Write(ctx, path, data)
	if err != nil {
		return "", exception.ErrInStorage.New(err)
	}

	return thumbnailURL, nil
}
