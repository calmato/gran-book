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

	v := md.Get(key)
	if len(v) == 0 {
		return "", xerrors.New("Metadata length is 0")
	}

	return v[0], nil
}
