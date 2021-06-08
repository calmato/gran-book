package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/golang/mock/gomock"
)

func TestChatDomainValidation_Room(t *testing.T) {
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
					ID: "00000000-0000-0000-0000-000000000000",
					UserIDs: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					},
					CreatedAt: current,
					UpdatedAt: current,
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

		t.Run(name, func(t *testing.T) {
			target := NewChatDomainValidation()

			got := target.Room(ctx, tc.args.room)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %#v, but %#v", tc.want, got)
			}
		})
	}
}
