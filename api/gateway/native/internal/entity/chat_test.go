package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/stretchr/testify/assert"
)

func TestChatRooms(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		rooms         ChatRooms
		expectUserIDs []string
	}{
		{
			name: "success",
			rooms: ChatRooms{
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
			assert.Equal(t, tt.expectUserIDs, tt.rooms.UserIDs())
		})
	}
}
