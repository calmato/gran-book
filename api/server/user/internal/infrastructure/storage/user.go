package storage

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	gcs "github.com/calmato/gran-book/api/server/user/pkg/firebase/storage"
)

type userUploader struct {
	storage *gcs.Storage
}

// NewUserUploader - UserUploaderの生成
func NewUserUploader(s *gcs.Storage) user.Uploader {
	return &userUploader{
		storage: s,
	}
}

func (s *userUploader) Thumbnail(ctx context.Context, userID string, data []byte) (string, error) {
	path := getUserThumbnailPath(userID)

	thumbnailURL, err := s.storage.Write(ctx, path, data)
	if err != nil {
		return "", exception.ErrorInStorage.New(err)
	}

	return thumbnailURL, nil
}
