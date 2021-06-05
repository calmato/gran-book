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
				LatestMassage: &chat.Message{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_chat.NewMockValidation(ctrl)

		urm := mock_chat.NewMockRepository(ctrl)
		urm.EXPECT().CreateRoom(ctx, tc.Room).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewChatService(uvm, urm)

			got := target.CreateRoom(ctx, tc.Room)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.Room.ID == "" {
				t.Fatal("User.ID must be not null")
				return
			}

			if tc.Room.CreatedAt.IsZero() {
				t.Fatal("User.CreatedAt must be not null")
				return
			}

			if tc.Room.UpdatedAt.IsZero() {
				t.Fatal("User.UpdatedAt must be not null")
				return
			}
		})
	}
}
