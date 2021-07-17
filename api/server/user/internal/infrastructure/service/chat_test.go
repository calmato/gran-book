package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	"github.com/golang/mock/gomock"
)

func TestChatService_ListRoom(t *testing.T) {
	type args struct {
		query *domain.ListQuery
		uid   string
	}
	type want struct {
		rooms []*chat.Room
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)
		crm.EXPECT().ListRoom(ctx, tc.args.query, tc.args.uid).Return(tc.want.rooms, tc.want.err)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			rooms, err := target.ListRoom(ctx, tc.args.query, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(rooms, tc.want.rooms) {
				t.Fatalf("want %#v, but %#v", tc.want.rooms, rooms)
				return
			}
		})
	}
}

func TestChatService_GetRoom(t *testing.T) {
	type args struct {
		roomID string
		uid    string
	}
	type want struct {
		room *chat.Room
		err  error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				roomID: "00000000-0000-0000-0000-000000000000",
				uid:    "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				room: &chat.Room{
					ID:      "00000000-0000-0000-0000-000000000000",
					UserIDs: []string{"00000000-0000-0000-0000-000000000000"},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)
		crm.EXPECT().GetRoom(ctx, tc.args.roomID).Return(tc.want.room, tc.want.err)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			room, err := target.GetRoom(ctx, tc.args.roomID)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(room, tc.want.room) {
				t.Fatalf("want %#v, but %#v", tc.want.room, room)
				return
			}
		})
	}
}

func TestChatService_CreateRoom(t *testing.T) {
	type args struct {
		room *chat.Room
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				room: &chat.Room{
					UserIDs:       []string{"00000000-0000-0000-0000-000000000000"},
					LatestMessage: &chat.Message{},
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)
		crm.EXPECT().CreateRoom(ctx, tc.args.room).Return(tc.want)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			got := target.CreateRoom(ctx, tc.args.room)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if got == nil {
				if tc.args.room.ID == "" {
					t.Fatal("Room.ID must be not null")
					return
				}

				if tc.args.room.CreatedAt.IsZero() {
					t.Fatal("Room.CreatedAt must be not null")
					return
				}

				if tc.args.room.UpdatedAt.IsZero() {
					t.Fatal("Room.UpdatedAt must be not null")
					return
				}
			}
		})
	}
}

func TestChatService_CreateMessage(t *testing.T) {
	type args struct {
		room    *chat.Room
		message *chat.Message
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				room: &chat.Room{
					UserIDs: []string{"00000000-0000-0000-0000-000000000000"},
				},
				message: &chat.Message{
					UserID: "00000000-0000-0000-0000-000000000000",
					Text:   "テストメッセージ",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)
		crm.EXPECT().UpdateRoom(ctx, tc.args.room).Return(tc.want)
		crm.EXPECT().CreateMessage(ctx, tc.args.room, tc.args.message).Return(tc.want)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			got := target.CreateMessage(ctx, tc.args.room, tc.args.message)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if got == nil {
				if tc.args.message.ID == "" {
					t.Fatal("Message.ID must be not null")
					return
				}

				if tc.args.message.CreatedAt.IsZero() {
					t.Fatal("Message.CreatedAt must be not null")
					return
				}

				if tc.args.room.LatestMessage == nil {
					t.Fatal("Room.LatestMessage must be not null")
					return
				}

				if tc.args.room.UpdatedAt.IsZero() {
					t.Fatal("Room.UpdatedAt must be not null")
					return
				}
			}
		})
	}
}

func TestChatService_UploadImage(t *testing.T) {
	type args struct {
		roomID string
		image  []byte
	}
	type want struct {
		imageURL string
		err      error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				roomID: "00000000-0000-0000-0000-000000000000",
				image:  []byte{},
			},
			want: want{
				imageURL: "https://calmato.com/chat_images",
				err:      nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)

		cum := mock_chat.NewMockUploader(ctrl)
		cum.EXPECT().Image(ctx, tc.args.roomID, tc.args.image).Return(tc.want.imageURL, tc.want.err)

		cmm := mock_chat.NewMockMessaging(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			imageURL, err := target.UploadImage(ctx, tc.args.roomID, tc.args.image)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(imageURL, tc.want.imageURL) {
				t.Fatalf("want %#v, but %#v", tc.want.imageURL, imageURL)
				return
			}
		})
	}
}

func TestChatService_PushCreateRoom(t *testing.T) {
	type args struct {
		room *chat.Room
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				room: &chat.Room{
					ID:          "000000000-0000-0000-0000-000000000000",
					UserIDs:     []string{"000000000-0000-0000-0000-000000000000"},
					CreatedAt:   current,
					UpdatedAt:   current,
					InstanceIDs: []string{"ExponentPushToken[XXXXXXXXXXXXXXXXXXXXXX]"},
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)
		cmm.EXPECT().PushCreateRoom(tc.args.room).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			got := target.PushCreateRoom(ctx, tc.args.room)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestChatService_PushNewMessage(t *testing.T) {
	type args struct {
		room    *chat.Room
		message *chat.Message
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				room: &chat.Room{
					ID:          "000000000-0000-0000-0000-000000000000",
					UserIDs:     []string{"000000000-0000-0000-0000-000000000000"},
					CreatedAt:   current,
					UpdatedAt:   current,
					InstanceIDs: []string{"ExponentPushToken[XXXXXXXXXXXXXXXXXXXXXX]"},
				},
				message: &chat.Message{
					ID:        "000000000-0000-0000-0000-000000000000",
					Text:      "テストメッセージ",
					UserID:    "000000000-0000-0000-0000-000000000000",
					Username:  "テストユーザ",
					CreatedAt: current,
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cvm := mock_chat.NewMockValidation(ctrl)

		crm := mock_chat.NewMockRepository(ctrl)

		cum := mock_chat.NewMockUploader(ctrl)

		cmm := mock_chat.NewMockMessaging(ctrl)
		cmm.EXPECT().PushNewMessage(tc.args.room, tc.args.message).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm, cum, cmm)

			got := target.PushNewMessage(ctx, tc.args.room, tc.args.message)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}
