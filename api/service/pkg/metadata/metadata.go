package metadata

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

var (
	ErrInvalidMetadata    = errors.New("metadata: this metadata is invalid")
	ErrNotFoundMetadata   = errors.New("metadata: this metadata is not found")
	ErrInvalidContext     = errors.New("metadata: this context has wrong type")
	ErrNotRetrieveContext = errors.New("metadata: this context could not retrieve")
)

// Get - メタデータの取得
func Get(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("%w, connot be retrieved from context", ErrInvalidMetadata)
	}

	v := md.Get(key)
	if len(v) == 0 {
		return "", ErrNotFoundMetadata
	}

	return v[0], nil
}
