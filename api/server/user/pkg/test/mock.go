package test

import (
	"errors"
	"testing"
	"time"

	mock_application "github.com/calmato/gran-book/api/server/user/mock/application"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/interface/validation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst)
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst)
)

type TestResponse struct {
	Code    codes.Code
	Message proto.Message
}

type Mocks struct {
	AdminRequestValidation *mock_validation.MockAdminRequestValidation
	AuthRequestValidation  *mock_validation.MockAuthRequestValidation
	ChatApplication        *mock_application.MockChatApplication
	ChatDomainValidation   *mock_chat.MockValidation
	ChatRepository         *mock_chat.MockRepository
	ChatRequestValidation  *mock_validation.MockChatRequestValidation
	ChatUploader           *mock_chat.MockUploader
	UserApplication        *mock_application.MockUserApplication
	UserDomainValidation   *mock_user.MockValidation
	UserRepository         *mock_user.MockRepository
	UserRequestValidation  *mock_validation.MockUserRequestValidation
	UserUploader           *mock_user.MockUploader
}

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
