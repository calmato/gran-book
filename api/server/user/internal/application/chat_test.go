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
	type args struct {
		uid   string
		input *input.CreateRoom
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
				uid: "00000000-0000-0000-0000-000000000000",
				input: &input.CreateRoom{
					UserIDs: []string{"00000000-0000-0000-0000-000000000000"},
				},
			},
			want: want{
				room: &chat.Room{
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

		crvm := mock_validation.NewMockChatRequestValidation(ctrl)
		crvm.EXPECT().CreateRoom(tc.args.input).Return(nil)

		csm := mock_chat.NewMockService(ctrl)
		csm.EXPECT().ValidationRoom(ctx, tc.want.room).Return(nil)
		csm.EXPECT().CreateRoom(ctx, tc.want.room).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewChatApplication(crvm, csm)

			room, err := target.CreateRoom(ctx, tc.args.input, tc.args.uid)
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
