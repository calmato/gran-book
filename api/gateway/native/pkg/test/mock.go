package test

import (
	"errors"
	"time"

	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/golang/mock/gomock"
)

const filename = "calmato.png"

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = datetime.FormatTime(time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst))
	DateMock = datetime.FormatDate(time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst))
)

type TestResponse struct {
	Code int
	Body interface{}
}

type Mocks struct {
	AdminService *mock_user.MockAdminServiceClient
	AuthService  *mock_user.MockAuthServiceClient
	BookService  *mock_book.MockBookServiceClient
	ChatService  *mock_chat.MockChatServiceClient
	UserService  *mock_user.MockUserServiceClient
}

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminService: mock_user.NewMockAdminServiceClient(ctrl),
		AuthService:  mock_user.NewMockAuthServiceClient(ctrl),
		BookService:  mock_book.NewMockBookServiceClient(ctrl),
		ChatService:  mock_chat.NewMockChatServiceClient(ctrl),
		UserService:  mock_user.NewMockUserServiceClient(ctrl),
	}
}
