package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type GrpcMetadata struct {
	Ctx context.Context
}

// Get - メタデータの取得
func (m *GrpcMetadata) Get(key string) (string, error) {
	md, ok := metadata.FromIncomingContext(m.Ctx)
	if !ok {
		return "", ErrInvalidMetadata
	}

	v := md.Get(key)
	if len(v) == 0 {
		return "", ErrNotFoundMetadata
	}

	return v[0], nil
}

// Set - メタデータの代入
func (m *GrpcMetadata) Set(key string, value string) {
	m.Ctx = metadata.AppendToOutgoingContext(m.Ctx, key, value)
}
