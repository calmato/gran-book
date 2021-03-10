package registry

import (
	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	gcs "github.com/calmato/gran-book/api/server/user/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct{}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(db *repository.Client, s *gcs.Storage) *Registry {
	return &Registry{}
}
