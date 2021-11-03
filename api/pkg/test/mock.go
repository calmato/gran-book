package test

import (
	"errors"
	"time"

	mock_book_application "github.com/calmato/gran-book/api/mock/book/application"
	mock_book "github.com/calmato/gran-book/api/mock/book/domain/book"
	mock_book_validation "github.com/calmato/gran-book/api/mock/book/interface/validation"
	mock_book_service "github.com/calmato/gran-book/api/mock/proto/book"
	mock_chat_service "github.com/calmato/gran-book/api/mock/proto/chat"
	mock_user_service "github.com/calmato/gran-book/api/mock/proto/user"
	mock_user_application "github.com/calmato/gran-book/api/mock/user/application"
	mock_chat "github.com/calmato/gran-book/api/mock/user/domain/chat"
	mock_user "github.com/calmato/gran-book/api/mock/user/domain/user"
	mock_user_validation "github.com/calmato/gran-book/api/mock/user/interface/validation"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

const filename = "calmato.png"

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local()
)

type HTTPResponse struct {
	Code int
	Body interface{}
}

type Response struct {
	Code    codes.Code
	Message proto.Message
}

type Mocks struct {
	AdminRequestValidation *mock_user_validation.MockAdminRequestValidation
	AdminService           *mock_user_service.MockAdminServiceClient
	AuthRequestValidation  *mock_user_validation.MockAuthRequestValidation
	AuthService            *mock_user_service.MockAuthServiceClient
	BookApplication        *mock_book_application.MockBookApplication
	BookDomainValidation   *mock_book.MockValidation
	BookRepository         *mock_book.MockRepository
	BookRequestValidation  *mock_book_validation.MockBookRequestValidation
	BookService            *mock_book_service.MockBookServiceClient
	ChatApplication        *mock_user_application.MockChatApplication
	ChatDomainValidation   *mock_chat.MockValidation
	ChatRepository         *mock_chat.MockRepository
	ChatRequestValidation  *mock_user_validation.MockChatRequestValidation
	ChatService            *mock_chat_service.MockChatServiceClient
	ChatUploader           *mock_chat.MockUploader
	UserApplication        *mock_user_application.MockUserApplication
	UserDomainValidation   *mock_user.MockValidation
	UserRepository         *mock_user.MockRepository
	UserRequestValidation  *mock_user_validation.MockUserRequestValidation
	UserService            *mock_user_service.MockUserServiceClient
	UserUploader           *mock_user.MockUploader
}

type DBMocks struct {
	UserDB        *database.Client
	BookDB        *database.Client
	InformationDB *database.Client
}

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminRequestValidation: mock_user_validation.NewMockAdminRequestValidation(ctrl),
		AdminService:           mock_user_service.NewMockAdminServiceClient(ctrl),
		AuthRequestValidation:  mock_user_validation.NewMockAuthRequestValidation(ctrl),
		AuthService:            mock_user_service.NewMockAuthServiceClient(ctrl),
		BookApplication:        mock_book_application.NewMockBookApplication(ctrl),
		BookDomainValidation:   mock_book.NewMockValidation(ctrl),
		BookRepository:         mock_book.NewMockRepository(ctrl),
		BookRequestValidation:  mock_book_validation.NewMockBookRequestValidation(ctrl),
		BookService:            mock_book_service.NewMockBookServiceClient(ctrl),
		ChatApplication:        mock_user_application.NewMockChatApplication(ctrl),
		ChatDomainValidation:   mock_chat.NewMockValidation(ctrl),
		ChatRepository:         mock_chat.NewMockRepository(ctrl),
		ChatRequestValidation:  mock_user_validation.NewMockChatRequestValidation(ctrl),
		ChatService:            mock_chat_service.NewMockChatServiceClient(ctrl),
		ChatUploader:           mock_chat.NewMockUploader(ctrl),
		UserApplication:        mock_user_application.NewMockUserApplication(ctrl),
		UserDomainValidation:   mock_user.NewMockValidation(ctrl),
		UserRepository:         mock_user.NewMockRepository(ctrl),
		UserRequestValidation:  mock_user_validation.NewMockUserRequestValidation(ctrl),
		UserService:            mock_user_service.NewMockUserServiceClient(ctrl),
		UserUploader:           mock_user.NewMockUploader(ctrl),
	}
}
