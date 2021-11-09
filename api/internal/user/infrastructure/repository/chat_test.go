package repository

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/pkg/firebase/firestore"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatRepository_ListRoom(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mocks, err := test.NewFirebaseMock(ctx)
	require.NoError(t, err)
	mocks.Firestore.Close()

	err = mocks.DeleteAll(ctx)
	require.NoError(t, err)

	rooms := make([]*chat.Room, 2)
	rooms[0] = testChatRoom("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000")
	rooms[1] = testChatRoom("11111111-1111-1111-1111-111111111111", "00000000-0000-0000-0000-000000000000")

	for i := range rooms {
		err = mocks.Firestore.Set(ctx, getChatRoomCollection(), rooms[i].ID, rooms[i])
		require.NoError(t, err)
	}

	type args struct {
		params  *firestore.Params
		queries []*firestore.Query
	}
	type want struct {
		rooms chat.Rooms
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, mocks *test.FirebaseMocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.FirebaseMocks) {},
			args: args{
				params: &firestore.Params{},
				queries: []*firestore.Query{
					{
						Field:    "users",
						Operator: "array-contains",
						Value:    "00000000-0000-0000-0000-000000000000",
					},
				},
			},
			want: want{
				rooms: rooms,
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			mocks, err := test.NewFirebaseMock(ctx)
			defer mocks.Firestore.Close()
			require.NoError(t, err)
			tt.setup(ctx, t, mocks)

			target := NewChatRepository(mocks.Firestore)
			crs, err := target.ListRoom(ctx, tt.args.params, tt.args.queries)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.rooms {
				assert.Contains(t, crs, tt.want.rooms[i])
			}
		})
	}
}

func TestChatRepository_GetRoom(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mocks, err := test.NewFirebaseMock(ctx)
	require.NoError(t, err)
	defer mocks.Firestore.Close()

	err = mocks.DeleteAll(ctx)
	require.NoError(t, err)

	room1 := testChatRoom("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000")
	err = mocks.Firestore.Set(ctx, getChatRoomCollection(), room1.ID, room1)
	require.NoError(t, err)

	type args struct {
		roomID string
	}
	type want struct {
		room  *chat.Room
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, mocks *test.FirebaseMocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.FirebaseMocks) {},
			args: args{
				roomID: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				room:  room1,
				isErr: false,
			},
		},
		{
			name:  "failed to not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.FirebaseMocks) {},
			args: args{
				roomID: "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				room:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			mocks, err := test.NewFirebaseMock(ctx)
			defer mocks.Firestore.Close()
			require.NoError(t, err)
			tt.setup(ctx, t, mocks)

			target := NewChatRepository(mocks.Firestore)
			cr, err := target.GetRoom(ctx, tt.args.roomID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.room, cr)
		})
	}
}

func testChatRoom(id string, userIDs ...string) *chat.Room {
	users := []string{
		"12345678-1234-1234-123456789012",
		"23456789-2345-2345-234567890123",
	}

	return &chat.Room{
		ID:        id,
		UserIDs:   append(users, userIDs...),
		CreatedAt: test.TimeMock.UTC(),
		UpdatedAt: test.TimeMock.UTC(),
	}
}
