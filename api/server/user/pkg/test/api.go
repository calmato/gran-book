package test

import (
	"testing"

	mock_application "github.com/calmato/gran-book/api/server/user/mock/application"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/interface/validation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminRequestValidation: mock_validation.NewMockAdminRequestValidation(ctrl),
		AuthRequestValidation:  mock_validation.NewMockAuthRequestValidation(ctrl),
		ChatApplication:        mock_application.NewMockChatApplication(ctrl),
		ChatDomainValidation:   mock_chat.NewMockValidation(ctrl),
		ChatRepository:         mock_chat.NewMockRepository(ctrl),
		ChatRequestValidation:  mock_validation.NewMockChatRequestValidation(ctrl),
		ChatUploader:           mock_chat.NewMockUploader(ctrl),
		UserApplication:        mock_application.NewMockUserApplication(ctrl),
		UserDomainValidation:   mock_user.NewMockValidation(ctrl),
		UserRepository:         mock_user.NewMockRepository(ctrl),
		UserRequestValidation:  mock_validation.NewMockUserRequestValidation(ctrl),
		UserUploader:           mock_user.NewMockUploader(ctrl),
	}
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
