package test

import (
	"errors"
	"time"

	mock_application "github.com/calmato/gran-book/api/server/user/mock/application"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"google.golang.org/grpc/codes"
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

type DBMocks struct {
	UserDB *database.Client
}
