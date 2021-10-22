package test

import (
	"errors"
	"time"

	mock_book_application "github.com/calmato/gran-book/api/service/mock/book/application"
	mock_book "github.com/calmato/gran-book/api/service/mock/book/domain/book"
	mock_book_validation "github.com/calmato/gran-book/api/service/mock/book/interface/validation"
	mock_user_application "github.com/calmato/gran-book/api/service/mock/user/application"
	mock_chat "github.com/calmato/gran-book/api/service/mock/user/domain/chat"
	mock_user "github.com/calmato/gran-book/api/service/mock/user/domain/user"
	mock_user_validation "github.com/calmato/gran-book/api/service/mock/user/interface/validation"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local()
)

type Response struct {
	Code    codes.Code
	Message proto.Message
}

type Mocks struct {
	AdminRequestValidation *mock_user_validation.MockAdminRequestValidation
	AuthRequestValidation  *mock_user_validation.MockAuthRequestValidation
	BookApplication        *mock_book_application.MockBookApplication
	BookDomainValidation   *mock_book.MockValidation
	BookRepository         *mock_book.MockRepository
	BookRequestValidation  *mock_book_validation.MockBookRequestValidation
	ChatApplication        *mock_user_application.MockChatApplication
	ChatDomainValidation   *mock_chat.MockValidation
	ChatRepository         *mock_chat.MockRepository
	ChatRequestValidation  *mock_user_validation.MockChatRequestValidation
	ChatUploader           *mock_chat.MockUploader
	UserApplication        *mock_user_application.MockUserApplication
	UserDomainValidation   *mock_user.MockValidation
	UserRepository         *mock_user.MockRepository
	UserRequestValidation  *mock_user_validation.MockUserRequestValidation
	UserUploader           *mock_user.MockUploader
}

type DBMocks struct {
	UserDB *database.Client
	BookDB *database.Client
}
