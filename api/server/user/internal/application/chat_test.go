package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_chat "github.com/calmato/gran-book/api/server/user/mock/domain/chat"
	"github.com/golang/mock/gomock"
)

func TestChatApplication_CreateRoom(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		Input    *input.CreateRoom
		Expected struct {
			Room  *chat.Room
			Error error
		}
	}{
		"ok": {
			UID: "00000000-0000-0000-0000-000000000000",
			Input: &input.CreateRoom{
				UserIDs: []string{"00000000-0000-0000-0000-000000000000"},
			},
			Expected: struct {
				Room  *chat.Room
				Error error
			}{
				Room: &chat.Room{
					UserIDs: []string{"00000000-0000-0000-0000-000000000000"},
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		crvm := mock_validation.NewMockChatRequestValidation(ctrl)
		crvm.EXPECT().CreateRoom(tc.Input).Return(nil)

		csm := mock_chat.NewMockService(ctrl)
		csm.EXPECT().CreateRoom(ctx, tc.Expected.Room).Return(tc.Expected.Error)
		csm.EXPECT().ValidationRoom(ctx, tc.Expected.Room).Return(nil)

		t.Run(result, func(t *testing.T) {
			target := NewChatApplication(crvm, csm)

			got, err := target.CreateRoom(ctx, tc.Input, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Room) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Room, got)
				return
			}
		})
	}
}
