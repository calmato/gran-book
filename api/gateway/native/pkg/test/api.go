package test

import (
	"github.com/calmato/gran-book/api/gateway/native/mock/book"
	"github.com/calmato/gran-book/api/gateway/native/mock/chat"
	"github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/golang/mock/gomock"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminService: user.NewMockAdminServiceClient(ctrl),
		AuthService:  user.NewMockAuthServiceClient(ctrl),
		BookService:  book.NewMockBookServiceClient(ctrl),
		ChatService:  chat.NewMockChatServiceClient(ctrl),
		UserService:  user.NewMockUserServiceClient(ctrl),
	}
}
