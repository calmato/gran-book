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
	current := time.Now()

	testCases := map[string]struct {
		Room     *chat.Room
		Expected error
	}{
		"ok": {
			Room: &chat.Room{
				ID:            "00000000-0000-0000-0000-000000000000",
				UserIDs:       []string{"00000000-0000-0000-0000-000000000000"},
				CreatedAt:     current,
				UpdatedAt:     current,
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

		t.Run(result, func(t *testing.T) {
			target := NewChatDomainValidation()

			got := target.Room(ctx, tc.Room)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}
