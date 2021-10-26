package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/stretchr/testify/assert"
)

func TestChatRoom(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		room   *chat.Room
		expect *ChatRoom
	}{
		{
			name: "success",
			room: &chat.Room{
				Id: "00000000-0000-0000-0000-000000000000",
				UserIds: []string{
					"12345678-1234-1234-1234-123456789012",
					"23456789-2345-2345-2345-234567890123",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &ChatRoom{
				Room: &chat.Room{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
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
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		rooms  []*chat.Room
		expect ChatRooms
	}{
		{
			name: "success",
			rooms: []*chat.Room{
				{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id: "11111111-1111-1111-1111-111111111111",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: ChatRooms{
				{
					Room: &chat.Room{
						Id: "00000000-0000-0000-0000-000000000000",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				{
					Room: &chat.Room{
						Id: "11111111-1111-1111-1111-111111111111",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRooms(tt.rooms)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChatRooms_UserIDs(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		rooms  []*chat.Room
		expect []string
	}{
		{
			name: "success",
			rooms: []*chat.Room{
				{
					Id: "00000000-0000-0000-0000-000000000000",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id: "11111111-1111-1111-1111-111111111111",
					UserIds: []string{
						"12345678-1234-1234-1234-123456789012",
						"23456789-2345-2345-2345-234567890123",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []string{
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
			assert.Equal(t, tt.expect, actual.UserIDs())
		})
	}
}

func TestChatMessage(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name    string
		message *chat.Message
		expect  *ChatMessage
	}{
		{
			name: "success",
			message: &chat.Message{
				Id:        "00000000-0000-0000-0000-000000000000",
				UserId:    "12345678-1234-1234-1234-123456789012",
				Text:      "テストメッセージです。",
				Image:     "",
				CreatedAt: now,
			},
			expect: &ChatMessage{
				Message: &chat.Message{
					Id:        "00000000-0000-0000-0000-000000000000",
					UserId:    "12345678-1234-1234-1234-123456789012",
					Text:      "テストメッセージです。",
					Image:     "",
					CreatedAt: now,
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
