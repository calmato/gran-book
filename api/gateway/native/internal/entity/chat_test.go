package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/stretchr/testify/assert"
)

func TestChatRoom(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		room   *pb.Room
		expect *ChatRoom
	}{
		{
			name: "success",
			room: &pb.Room{
				Id: "00000000-0000-0000-0000-000000000000",
				UserIds: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
			expect: &ChatRoom{
				Room: &pb.Room{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRoom(tt.room)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChatRooms(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		rooms         []*pb.Room
		expect        ChatRooms
		expectUserIDs []string
	}{
		{
			name: "success",
			rooms: []*pb.Room{
				{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
				{
					Id: "11111111-1111-1111-1111-111111111111",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			expect: ChatRooms{
				{
					Room: &pb.Room{
						Id: "00000000-0000-0000-0000-000000000000",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
				{
					Room: &pb.Room{
						Id: "11111111-1111-1111-1111-111111111111",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
			},
			expectUserIDs: []string{
				"12345678-1234-1234-1234-123456789012",
				"23456789-2345-2345-2345-234567890123",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRooms(tt.rooms)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectUserIDs, actual.UserIDs())
		})
	}
}

func TestChatMessage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		message *pb.Message
		expect  *ChatMessage
	}{
		{
			name: "success",
			message: &pb.Message{
				Id:        "00000000-0000-0000-0000-000000000000",
				UserId:    "12345678-1234-1234-1234-123456789012",
				Text:      "テストメッセージです。",
				Image:     "",
				CreatedAt: test.TimeMock,
			},
			expect: &ChatMessage{
				Message: &pb.Message{
					Id:        "00000000-0000-0000-0000-000000000000",
					UserId:    "12345678-1234-1234-1234-123456789012",
					Text:      "テストメッセージです。",
					Image:     "",
					CreatedAt: test.TimeMock,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatMessage(tt.message)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
