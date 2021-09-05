package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
)

func TestAuthServer_ListRoom(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")
	room2 := testChatRoom("room02")

	type args struct {
		req *pb.ListChatRoomRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListChatRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					ListRoom(ctx, "user01", gomock.Any()).
					Return(chat.Rooms{room1, room2}, nil)
			},
			args: args{
				req: &pb.ListChatRoomRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getChatRoomListResponse([]*chat.Room{room1, room2}),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListChatRoom(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListChatRoomRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListChatRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					ListRoom(ctx, "user01", gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListChatRoomRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatServer(mocks.ChatRequestValidation, mocks.ChatApplication)

			res, err := target.ListRoom(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_CreateRoom(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.CreateChatRoomRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					CreateRoom(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				req: &pb.CreateChatRoomRequest{
					UserIds: []string{"user01", "user02"},
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.ChatRoomResponse{
					Room: &pb.ChatRoom{
						UserIds:       []string{"user01", "user02"},
						CreatedAt:     "",
						UpdatedAt:     "",
						LatestMessage: &pb.ChatMessage{},
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatRoom(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateChatRoomRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					CreateRoom(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateChatRoomRequest{
					UserIds: []string{"user01", "user02"},
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatServer(mocks.ChatRequestValidation, mocks.ChatApplication)

			res, err := target.CreateRoom(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_CreateMessage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")

	type args struct {
		req *pb.CreateChatMessageRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(room1, nil)
				mocks.ChatApplication.EXPECT().
					CreateMessage(ctx, room1, gomock.Any()).
					Return(nil)
			},
			args: args{
				req: &pb.CreateChatMessageRequest{
					UserId: "user01",
					RoomId: "room01",
					Text:   "テストメッセージ",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.ChatMessageResponse{
					Message: &pb.ChatMessage{
						UserId:    "user01",
						Text:      "テストメッセージ",
						CreatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatMessage(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateChatMessageRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateChatMessageRequest{
					UserId: "user01",
					RoomId: "room01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateChatMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(room1, nil)
				mocks.ChatApplication.EXPECT().
					CreateMessage(ctx, room1, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateChatMessageRequest{
					UserId: "user01",
					RoomId: "room01",
					Text:   "テストメッセージ",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatServer(mocks.ChatRequestValidation, mocks.ChatApplication)

			res, err := target.CreateMessage(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func testChatRoom(id string) *chat.Room {
	return &chat.Room{
		ID: id,
		UserIDs: []string{
			"12345678-1234-1234-123456789012",
			"23456789-2345-2345-234567890123",
		},
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}

func testChatMessage(id string) *chat.Message {
	return &chat.Message{
		ID:        id,
		Text:      "テストメッセージです",
		Image:     "",
		UserID:    "12345678-1234-1234-123456789012",
		CreatedAt: test.TimeMock,
	}
}
