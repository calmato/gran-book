package registry

import (
	"github.com/calmato/gran-book/api/server/notification/lib/firebase/firestore"
	gcs "github.com/calmato/gran-book/api/server/notification/lib/firebase/storage"
)

// Registry - DIコンテナ
type Registry struct{}

// NewRegistry - internalディレクトリ配下のファイルを読み込み
func NewRegistry(fs *firestore.Client, s *gcs.Storage) *Registry {
	return &Registry{}
}
