package service

import (
	"context"
	"reflect"
	"testing"

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

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm)

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

		t.Run(name, func(t *testing.T) {
			target := NewChatService(cvm, crm)

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
