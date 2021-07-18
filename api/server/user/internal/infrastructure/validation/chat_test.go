package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/golang/mock/gomock"
	"golang.org/x/xerrors"
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

func TestChatDomainValidation_Message(t *testing.T) {
	type args struct {
		message *chat.Message
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				message: &chat.Message{
					ID:        "00000000-0000-0000-0000-000000000000",
					UserID:    "00000000-0000-0000-0000-000000000000",
					Text:      "テストメッセージ",
					Image:     "",
					CreatedAt: current,
				},
			},
			want: nil,
		},
		"ng_requires_either_text_or_image": {
			args: args{
				message: &chat.Message{
					ID:        "00000000-0000-0000-0000-000000000000",
					UserID:    "00000000-0000-0000-0000-000000000000",
					Text:      "",
					Image:     "",
					CreatedAt: current,
				},
			},
			want: exception.InvalidDomainValidation.New(
				xerrors.New("This message requires either text or image."),
				&exception.ValidationError{
					Field:   "text",
					Message: exception.RequiredMessage,
				},
			),
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewChatDomainValidation()

			got := target.Message(ctx, tc.args.message)
			if tc.want == nil {
				if got != nil {
					t.Errorf("want %#v, but %#v", tc.want, got)
					return
				}
			} else {
				if got == nil {
					t.Errorf("want %#v, but %#v", tc.want, got)
					return
				}
			}
		})
	}
}
