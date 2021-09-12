package test

import (
	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/golang/mock/gomock"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminService: mock_user.NewMockAdminServiceClient(ctrl),
		AuthService:  mock_user.NewMockAuthServiceClient(ctrl),
		BookService:  mock_book.NewMockBookServiceClient(ctrl),
		ChatService:  mock_chat.NewMockChatServiceClient(ctrl),
		UserService:  mock_user.NewMockUserServiceClient(ctrl),
	}
}
