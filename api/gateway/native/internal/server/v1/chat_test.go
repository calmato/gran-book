package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/golang/mock/gomock"
)

func TestChat_ListChatRoom(t *testing.T) {
	t.Parallel()

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	rooms := make([]*chat.Room, 2)
	rooms[0] = testChatRoom("00000000-0000-0000-0000-000000000000")
	rooms[1] = testChatRoom("11111111-1111-1111-1111-111111111111")
	users := make([]*user.User, 2)
	users[0] = testUser("12345678-1234-1234-123456789012")
	users[1] = testUser("23456789-2345-2345-234567890123")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().
					ListRoom(gomock.Any(), &chat.ListRoomRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: "0",
					}).
					Return(&chat.RoomListResponse{Rooms: rooms}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-123456789012",
							"23456789-2345-2345-234567890123",
						},
					}).
					Return(&user.UserListResponse{
						Users:  users,
						Total:  2,
						Limit:  2,
						Offset: 0,
					}, nil)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewChatRoomListResponse(
					entity.NewChatRooms(rooms),
					entity.NewUsers(users).Map(),
				),
			},
		},
		{
			name:  "failed to invalid limit query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?limit=1.1",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?offset=0.1",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list room",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().
					ListRoom(gomock.Any(), &chat.ListRoomRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: "0",
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().
					ListRoom(gomock.Any(), &chat.ListRoomRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: "0",
					}).
					Return(&chat.RoomListResponse{Rooms: rooms}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-123456789012",
							"23456789-2345-2345-234567890123",
						},
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/chat" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestChat_CreateChatRoom(t *testing.T) {
	t.Parallel()

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	room := testChatRoom("00000000-0000-0000-0000-000000000000")
	users := make([]*user.User, 3)
	users[0] = testUser("12345678-1234-1234-1234-123456789012")
	users[1] = testUser("23456789-2345-2345-2345-234567890123")
	users[2] = testUser("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.CreateChatRoomRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(&user.UserListResponse{
						Users:  users,
						Total:  3,
						Limit:  3,
						Offset: 0,
					}, nil)
				mocks.ChatService.EXPECT().
					CreateRoom(gomock.Any(), &chat.CreateRoomRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(&chat.RoomResponse{Room: room}, nil)
			},
			req: &request.CreateChatRoomRequest{
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewChatRoomResponse(
					entity.NewChatRoom(room),
					entity.NewUsers(users).Map(),
				),
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.CreateChatRoomRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateChatRoomRequest{
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateChatRoomRequest{
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to match user ids",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(&user.UserListResponse{
						Users:  users[:2],
						Total:  2,
						Limit:  3,
						Offset: 0,
					}, nil)
			},
			req: &request.CreateChatRoomRequest{
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
			},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create room",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(&user.UserListResponse{
						Users:  users,
						Total:  3,
						Limit:  3,
						Offset: 0,
					}, nil)
				mocks.ChatService.EXPECT().
					CreateRoom(gomock.Any(), &chat.CreateRoomRequest{
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
							"00000000-0000-0000-0000-000000000000",
						},
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateChatRoomRequest{
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/chat"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestChat_CreateChatTextMessage(t *testing.T) {
	t.Parallel()

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	message1 := testChatMessage("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.CreateChatMessageRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().
					CreateMessage(gomock.Any(), &chat.CreateMessageRequest{
						RoomId: "00000000-0000-0000-0000-000000000000",
						UserId: "00000000-0000-0000-0000-000000000000",
						Text:   "テストメッセージです。",
					}).
					Return(&chat.MessageResponse{Message: message1}, nil)
			},
			req: &request.CreateChatMessageRequest{
				Text: "テストメッセージです。",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewChatMessageResponse(
					entity.NewChatMessage(message1),
					entity.NewAuth(auth),
				),
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.CreateChatMessageRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateChatMessageRequest{
				Text: "テストメッセージです。",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to create message",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().
					CreateMessage(gomock.Any(), &chat.CreateMessageRequest{
						RoomId: "00000000-0000-0000-0000-000000000000",
						UserId: "00000000-0000-0000-0000-000000000000",
						Text:   "テストメッセージです。",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateChatMessageRequest{
				Text: "テストメッセージです。",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/chat/00000000-0000-0000-0000-000000000000/messages/text"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestChat_CreateChatImageMessage(t *testing.T) {
	t.Parallel()

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	message1 := testChatMessage("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		field  string
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				client := mock_chat.NewMockChatService_UploadImageClient(ctrl)
				mocks.ChatService.EXPECT().UploadImage(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).AnyTimes().Return(nil)
				client.EXPECT().CloseAndRecv().Return(&chat.MessageResponse{
					Message: message1,
				}, nil)
			},
			field: "image",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewChatMessageResponse(
					entity.NewChatMessage(message1),
					entity.NewAuth(auth),
				),
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().GetAuth(gomock.Any(), &user.Empty{}).Return(nil, test.ErrMock)
			},
			field: "image",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
			},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create gRPC stream client",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.ChatService.EXPECT().UploadImage(gomock.Any()).Return(nil, test.ErrMock)
			},
			field: "image",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to send stream request",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				client := mock_chat.NewMockChatService_UploadImageClient(ctrl)
				mocks.ChatService.EXPECT().UploadImage(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).Return(test.ErrMock)
			},
			field: "image",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to upload image",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				client := mock_chat.NewMockChatService_UploadImageClient(ctrl)
				mocks.ChatService.EXPECT().UploadImage(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).AnyTimes().Return(nil)
				client.EXPECT().CloseAndRecv().Return(nil, test.ErrMock)
			},
			field: "image",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			path := "/v1/users/00000000-0000-0000-0000-000000000000/chat/00000000-0000-0000-0000-000000000000/messages/image"
			req := test.NewMultipartRequest(t, http.MethodPost, path, tt.field)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testChatRoom(id string) *chat.Room {
	return &chat.Room{
		Id: id,
		UserIds: []string{
			"12345678-1234-1234-123456789012",
			"23456789-2345-2345-234567890123",
		},
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}

func testChatMessage(id string) *chat.Message {
	return &chat.Message{
		Id:        id,
		Text:      "テストメッセージです",
		Image:     "",
		UserId:    "12345678-1234-1234-123456789012",
		CreatedAt: test.TimeMock,
	}
}