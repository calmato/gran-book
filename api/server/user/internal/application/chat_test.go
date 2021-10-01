package application

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestChatApplication_ListRoom(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")
	room2 := testChatRoom("room02")

	type args struct {
		userID string
		params *firestore.Params
	}
	type want struct {
		rooms chat.Rooms
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRepository.EXPECT().
					ListRoom(ctx, gomock.Any(), gomock.Any()).
					Return([]*chat.Room{room1, room2}, nil)
			},
			args: args{
				userID: "user01",
				params: &firestore.Params{},
			},
			want: want{
				rooms: chat.Rooms{room1, room2},
				err:   nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatApplication(
				mocks.ChatDomainValidation,
				mocks.ChatRepository,
				mocks.ChatUploader,
			)

			crs, err := target.ListRoom(ctx, tt.args.userID, tt.args.params)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.rooms, crs)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.rooms, crs)
		})
	}
}

func TestChatApplication_GetRoom(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")
	room1.UserIDs = append(room1.UserIDs, "user01")

	type args struct {
		roomID string
		userID string
	}
	type want struct {
		room *chat.Room
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRepository.EXPECT().
					GetRoom(ctx, "room01").
					Return(room1, nil)
			},
			args: args{
				roomID: "room01",
				userID: "user01",
			},
			want: want{
				room: room1,
				err:  nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRepository.EXPECT().
					GetRoom(ctx, "room01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				roomID: "room01",
				userID: "user01",
			},
			want: want{
				room: nil,
				err:  exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: forbidden",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatRepository.EXPECT().
					GetRoom(ctx, "room01").
					Return(&chat.Room{ID: "room01"}, nil)
			},
			args: args{
				roomID: "room01",
				userID: "user01",
			},
			want: want{
				room: nil,
				err:  exception.Forbidden.New(errNotJoinUserInRoom),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatApplication(
				mocks.ChatDomainValidation,
				mocks.ChatRepository,
				mocks.ChatUploader,
			)

			cr, err := target.GetRoom(ctx, tt.args.roomID, tt.args.userID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.room, cr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.room, cr)
		})
	}
}

func TestChatApplication_CreateRoom(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		room *chat.Room
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatDomainValidation.EXPECT().
					Room(ctx, gomock.Any()).
					Return(nil)
				mocks.ChatRepository.EXPECT().
					CreateRoom(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				room: &chat.Room{
					UserIDs: []string{"user01"},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: domain validation",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatDomainValidation.EXPECT().
					Room(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				room: &chat.Room{
					UserIDs: []string{"user01"},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatApplication(
				mocks.ChatDomainValidation,
				mocks.ChatRepository,
				mocks.ChatUploader,
			)

			err := target.CreateRoom(ctx, tt.args.room)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotEqual(t, tt.args.room.ID, "")
			require.NotZero(t, tt.args.room.CreatedAt)
			require.NotZero(t, tt.args.room.UpdatedAt)
		})
	}
}

func TestChatApplication_CreateMessage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")

	type args struct {
		room    *chat.Room
		message *chat.Message
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatDomainValidation.EXPECT().
					Message(ctx, gomock.Any()).
					Return(nil)
				mocks.ChatRepository.EXPECT().
					CreateMessage(ctx, "room01", gomock.Any()).
					Return(nil)
				mocks.ChatRepository.EXPECT().
					UpdateRoom(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				room: room1,
				message: &chat.Message{
					Text: "テストメッセージ",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: domain validation",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatDomainValidation.EXPECT().
					Message(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				room: room1,
				message: &chat.Message{
					Text: "テストメッセージ",
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatDomainValidation.EXPECT().
					Message(ctx, gomock.Any()).
					Return(nil)
				mocks.ChatRepository.EXPECT().
					CreateMessage(ctx, "room01", gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				room: room1,
				message: &chat.Message{
					Text: "テストメッセージ",
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatApplication(
				mocks.ChatDomainValidation,
				mocks.ChatRepository,
				mocks.ChatUploader,
			)

			err := target.CreateMessage(ctx, tt.args.room, tt.args.message)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotNil(t, tt.args.room.LatestMessage)
			require.NotZero(t, tt.args.room.UpdatedAt)
			require.NotEqual(t, tt.args.message.ID, "")
			require.NotZero(t, tt.args.message.CreatedAt)
		})
	}
}

func TestChatApplication_UploadImage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room1 := testChatRoom("room01")

	type args struct {
		room  *chat.Room
		image []byte
	}
	type want struct {
		imageURL string
		err      error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "succeess",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.ChatUploader.EXPECT().
					Image(ctx, "room01", []byte{}).
					Return("https://go.dev/images/gophers/ladder.svg", nil)
			},
			args: args{
				room:  room1,
				image: []byte{},
			},
			want: want{
				imageURL: "https://go.dev/images/gophers/ladder.svg",
				err:      nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatApplication(
				mocks.ChatDomainValidation,
				mocks.ChatRepository,
				mocks.ChatUploader,
			)

			imageURL, err := target.UploadImage(ctx, tt.args.room, tt.args.image)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.imageURL, imageURL)
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
