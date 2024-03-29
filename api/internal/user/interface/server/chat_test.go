package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/pkg/test"
	pb "github.com/calmato/gran-book/api/proto/chat"
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
		req *pb.ListRoomRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					ListRoom(ctx, "user01", gomock.Any()).
					Return(chat.Rooms{room1, room2}, nil)
			},
			args: args{
				req: &pb.ListRoomRequest{
					UserId: "user01",
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getChatRoomListResponse([]*chat.Room{room1, room2}),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListRoom(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListRoomRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					ListRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					ListRoom(ctx, "user01", gomock.Any()).
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListRoomRequest{
					UserId: "user01",
				},
			},
			want: &test.Response{
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
			test.GRPC(t, tt.want, res, err)
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
		req *pb.CreateRoomRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					CreateRoom(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				req: &pb.CreateRoomRequest{
					UserIds: []string{"user01", "user02"},
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.RoomResponse{
					Room: &pb.Room{
						UserIds:       []string{"user01", "user02"},
						CreatedAt:     "",
						UpdatedAt:     "",
						LatestMessage: nil,
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateRoom(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateRoomRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateRoom(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					CreateRoom(ctx, gomock.Any()).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateRoomRequest{
					UserIds: []string{"user01", "user02"},
				},
			},
			want: &test.Response{
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
			test.GRPC(t, tt.want, res, err)
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
		req *pb.CreateMessageRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(room1, nil)
				mocks.ChatApplication.EXPECT().
					CreateMessage(ctx, room1, gomock.Any()).
					Return(nil)
			},
			args: args{
				req: &pb.CreateMessageRequest{
					UserId: "user01",
					RoomId: "room01",
					Text:   "テストメッセージ",
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.MessageResponse{
					Message: &pb.Message{
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
					CreateMessage(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateMessageRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateMessageRequest{
					UserId: "user01",
					RoomId: "room01",
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRequestValidation.EXPECT().
					CreateMessage(gomock.Any()).
					Return(nil)
				mocks.ChatApplication.EXPECT().
					GetRoom(ctx, "room01", "user01").
					Return(room1, nil)
				mocks.ChatApplication.EXPECT().
					CreateMessage(ctx, room1, gomock.Any()).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateMessageRequest{
					UserId: "user01",
					RoomId: "room01",
					Text:   "テストメッセージ",
				},
			},
			want: &test.Response{
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
			test.GRPC(t, tt.want, res, err)
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
