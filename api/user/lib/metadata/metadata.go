package metadata

import (
	"context"

	"golang.org/x/xerrors"
	gmd "google.golang.org/grpc/metadata"
)

// Get - メタデータの取得
func Get(ctx context.Context, key string) (string, error) {
	md, ok := gmd.FromIncomingContext(ctx)
	if !ok {
		return "", xerrors.New("Metadata connot be retrieved from context")
	}

	return md.Get(key)[0], nil
}
