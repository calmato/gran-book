package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	"github.com/golang/mock/gomock"
)

func TestChatService_CreateRoom(t *testing.T) {
	testCases := map[string]struct {
		Room     *chat.Room
		Expected error
	}{
		"ok": {
			Room: &chat.Room{
				UserIDs:       []string{"00000000-0000-0000-0000-000000000000"},
				LatestMessage: &chat.Message{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rvm := mock_chat.NewMockValidation(ctrl)

		rrm := mock_chat.NewMockRepository(ctrl)
		rrm.EXPECT().CreateRoom(ctx, tc.Room).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewChatService(rvm, rrm)

			got := target.CreateRoom(ctx, tc.Room)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.Room.ID == "" {
				t.Fatal("Room.ID must be not null")
				return
			}

			if tc.Room.CreatedAt.IsZero() {
				t.Fatal("Room.CreatedAt must be not null")
				return
			}

			if tc.Room.UpdatedAt.IsZero() {
				t.Fatal("Room.UpdatedAt must be not null")
				return
			}
		})
	}
}
