package test

import (
	"github.com/calmato/gran-book/api/gateway/native/mock"
	"github.com/golang/mock/gomock"
)

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminService: mock.NewMockAdminServiceClient(ctrl),
		AuthService:  mock.NewMockAuthServiceClient(ctrl),
		BookService:  mock.NewMockBookServiceClient(ctrl),
		ChatService:  mock.NewMockChatServiceClient(ctrl),
		UserService:  mock.NewMockUserServiceClient(ctrl),
	}
}
