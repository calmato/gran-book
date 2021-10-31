package metadata

import (
	"context"
	"errors"
	"fmt"

	gmd "google.golang.org/grpc/metadata"
)

var errInvalidMetadata = errors.New("metadata: this metadata is invalid")

// Get - メタデータの取得
func Get(ctx context.Context, key string) (string, error) {
	md, ok := gmd.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("%w, connot be retrieved from context", errInvalidMetadata)
	}

	v := md.Get(key)
	if len(v) == 0 {
		return "", fmt.Errorf("%w, length is 0", errInvalidMetadata)
	}

	return v[0], nil
}
