package chat

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto/chat"
	"github.com/stretchr/testify/assert"
)

func TestRoom(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		room        *Room
		expectProto *pb.Room
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
			expectProto: &pb.Room{
				Id: "00000000-0000-0000-0000-000000000000",
				UserIds: []string{
					"12345678-1234-1234-123456789012",
					"23456789-2345-2345-234567890123",
				},
				CreatedAt: datetime.TimeToString(now),
				UpdatedAt: datetime.TimeToString(now),
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
		expectProto []*pb.Room
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
			expectProto: []*pb.Room{
				{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: datetime.TimeToString(now),
					UpdatedAt: datetime.TimeToString(now),
				},
				{
					Id: "11111111-1111-1111-1111-111111111111",
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
					CreatedAt: datetime.TimeToString(now),
					UpdatedAt: datetime.TimeToString(now),
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
		expectProto *pb.Message
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
			expectProto: &pb.Message{
				Id:        "00000000-0000-0000-0000-000000000000",
				Text:      "テストメッセージです",
				Image:     "",
				UserId:    "12345678-1234-1234-123456789012",
				CreatedAt: datetime.TimeToString(now),
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
