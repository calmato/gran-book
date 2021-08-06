package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAuthServer_GetAuth(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := testUser("user01")

	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		want  test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(user, nil)
			},
			want: test.TestResponse{
				Code:    codes.OK,
				Message: getAuthResponse(user),
			},
		},
		{
			name: "unauthorized",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			want: test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.GetAuth(ctx, &pb.Empty{})
			if tt.want.Code != codes.OK {
				require.Error(t, err)

				status, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.want.Code, status.Code(), status.Code().String())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want.Message, res)
		})
	}
}
