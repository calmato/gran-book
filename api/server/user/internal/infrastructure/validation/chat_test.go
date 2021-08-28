package validation

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestChatDomainValidation_Room(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		room *chat.Room
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  error
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				room: &chat.Room{
					ID: "00000000-0000-0000-0000-000000000000",
					UserIDs: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					},
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			want: nil,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatDomainValidation()

			err := target.Room(ctx, tt.args.room)
			if tt.want != nil {
				require.Equal(t, tt.want.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestChatDomainValidation_Message(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		message *chat.Message
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  error
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				message: &chat.Message{
					ID:        "00000000-0000-0000-0000-000000000000",
					UserID:    "00000000-0000-0000-0000-000000000000",
					Text:      "テストメッセージ",
					Image:     "",
					CreatedAt: test.TimeMock,
				},
			},
			want: nil,
		},
		{
			name:  "failed: require either text or image",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				message: &chat.Message{
					ID:        "00000000-0000-0000-0000-000000000000",
					UserID:    "00000000-0000-0000-0000-000000000000",
					Text:      "",
					Image:     "",
					CreatedAt: test.TimeMock,
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewChatDomainValidation()

			err := target.Message(ctx, tt.args.message)
			if tt.want != nil {
				require.Equal(t, tt.want.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
