package chat

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/proto/chat"
	"github.com/stretchr/testify/assert"
)

func TestRoom(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		room        *Room
		expectProto *chat.Room
	}{
		{
			name: "success",
			room: &Room{
				ID: "00000000-0000-0000-0000-000000000000",
				UserIDs: []string{
					"12345678-1234-1234-123456789012",
					"23456789-2345-2345-234567890123",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectProto: &chat.Room{
				Id: "00000000-0000-0000-0000-000000000000",
				UserIds: []string{
					"12345678-1234-1234-123456789012",
					"23456789-2345-2345-234567890123",
				},
				CreatedAt: datetime.FormatTime(now),
				UpdatedAt: datetime.FormatTime(now),
			},
		},
		{
			name: "success with messagee",
			room: &Room{
				ID: "00000000-0000-0000-0000-000000000000",
				UserIDs: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
				LatestMessage: &Message{
					ID:        "00000000-0000-0000-0000-000000000000",
					Text:      "テストメッセージです。",
					Image:     "",
					UserID:    "12345678-1234-1234-1234-123456789012",
					CreatedAt: now,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectProto: &chat.Room{
				Id: "00000000-0000-0000-0000-000000000000",
				UserIds: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
				LatestMessage: &chat.Message{
					Id:        "00000000-0000-0000-0000-000000000000",
					Text:      "テストメッセージです。",
					Image:     "",
					UserId:    "12345678-1234-1234-1234-123456789012",
					CreatedAt: datetime.FormatTime(now),
				},
				CreatedAt: datetime.FormatTime(now),
				UpdatedAt: datetime.FormatTime(now),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.room.Proto())
		})
	}
}

func TestRooms(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		rooms       Rooms
		expectProto []*chat.Room
	}{
		{
			name: "success",
			rooms: Rooms{
				{
					ID: "00000000-0000-0000-0000-000000000000",
					UserIDs: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID: "11111111-1111-1111-1111-111111111111",
					UserIDs: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expectProto: []*chat.Room{
				{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: datetime.FormatTime(now),
					UpdatedAt: datetime.FormatTime(now),
				},
				{
					Id: "11111111-1111-1111-1111-111111111111",
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: datetime.FormatTime(now),
					UpdatedAt: datetime.FormatTime(now),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.rooms.Proto())
		})
	}
}

func TestMessage(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		message     *Message
		expectProto *chat.Message
	}{
		{
			name: "success",
			message: &Message{
				ID:        "00000000-0000-0000-0000-000000000000",
				Text:      "テストメッセージです",
				Image:     "",
				UserID:    "12345678-1234-1234-123456789012",
				CreatedAt: now,
			},
			expectProto: &chat.Message{
				Id:        "00000000-0000-0000-0000-000000000000",
				Text:      "テストメッセージです",
				Image:     "",
				UserId:    "12345678-1234-1234-123456789012",
				CreatedAt: datetime.FormatTime(now),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.message.Proto())
		})
	}
}
