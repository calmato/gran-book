package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{}
}

func TestGRPC(t *testing.T, expect *TestResponse, res interface{}, err error) {
	if expect.Code != codes.OK {
		require.Error(t, err)

		status, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, expect.Code.String(), status.Code().String())
		return
	}

	require.NoError(t, err)
	require.Equal(t, expect.Message, res)
}
