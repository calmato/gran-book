package test

import (
	"testing"

	mock_application "github.com/calmato/gran-book/api/service/mock/book/application"
	mock_book "github.com/calmato/gran-book/api/service/mock/book/domain/book"
	mock_validation "github.com/calmato/gran-book/api/service/mock/book/interface/validation"
	mock_user_application "github.com/calmato/gran-book/api/service/mock/user/application"
	mock_chat "github.com/calmato/gran-book/api/service/mock/user/domain/chat"
	mock_user "github.com/calmato/gran-book/api/service/mock/user/domain/user"
	mock_user_validation "github.com/calmato/gran-book/api/service/mock/user/interface/validation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminRequestValidation: mock_user_validation.NewMockAdminRequestValidation(ctrl),
		AuthRequestValidation:  mock_user_validation.NewMockAuthRequestValidation(ctrl),
		BookApplication:        mock_application.NewMockBookApplication(ctrl),
		BookDomainValidation:   mock_book.NewMockValidation(ctrl),
		BookRepository:         mock_book.NewMockRepository(ctrl),
		BookRequestValidation:  mock_validation.NewMockBookRequestValidation(ctrl),
		ChatApplication:        mock_user_application.NewMockChatApplication(ctrl),
		ChatDomainValidation:   mock_chat.NewMockValidation(ctrl),
		ChatRepository:         mock_chat.NewMockRepository(ctrl),
		ChatRequestValidation:  mock_user_validation.NewMockChatRequestValidation(ctrl),
		ChatUploader:           mock_chat.NewMockUploader(ctrl),
		UserApplication:        mock_user_application.NewMockUserApplication(ctrl),
		UserDomainValidation:   mock_user.NewMockValidation(ctrl),
		UserRepository:         mock_user.NewMockRepository(ctrl),
		UserRequestValidation:  mock_user_validation.NewMockUserRequestValidation(ctrl),
		UserUploader:           mock_user.NewMockUploader(ctrl),
	}
}

func GRPC(t *testing.T, expect *Response, res interface{}, err error) {
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
