package metadata

import (
	"context"
	"fmt"

	gmd "google.golang.org/grpc/metadata"
)

type GrpcMetadata struct {
	Ctx context.Context
}

// Get - メタデータの取得
func (m *GrpcMetadata) Get(key string) (string, error) {
	md, ok := gmd.FromIncomingContext(m.Ctx)
	if !ok {
		return "", fmt.Errorf("metadata connot be retrieved from context")
	}

	v := md.Get(key)
	if len(v) == 0 {
		return "", fmt.Errorf("metadata length is 0")
	}

	return v[0], nil
}

// Set - メタデータの代入
func (m *GrpcMetadata) Set(key string, value string) {
	m.Ctx = gmd.AppendToOutgoingContext(m.Ctx, key, value)
}
