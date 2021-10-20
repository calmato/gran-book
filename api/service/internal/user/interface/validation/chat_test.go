package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/service/proto/chat"
	"github.com/stretchr/testify/assert"
)

func TestChatRequestValidation_ListRoom(t *testing.T) {
	t.Parallel()
	type args struct {
		req *chat.ListRoomRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &chat.ListRoomRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  200,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &chat.ListRoomRequest{
					UserId: "",
					Limit:  200,
					Offset: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &chat.ListRoomRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  201,
					Offset: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewChatRequestValidation()

			got := target.ListRoom(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestChatRequestValidation_CreateRoom(t *testing.T) {
	t.Parallel()
	type args struct {
		req *chat.CreateRoomRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &chat.CreateRoomRequest{
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
					},
				},
			},
			want: true,
		},
		{
			name: "validation error: UserIds.min_items",
			args: args{
				req: &chat.CreateRoomRequest{
					UserIds: []string{
						"12345678-1234-1234-123456789012",
					},
				},
			},
			want: false,
		},
		{
			name: "validation error: UserIds.max_items",
			args: args{
				req: &chat.CreateRoomRequest{
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"23456789-2345-2345-234567890123",
						"34567890-3456-3456-345678901234",
					},
				},
			},
			want: false,
		},
		{
			name: "validation error: UserIds.unique",
			args: args{
				req: &chat.CreateRoomRequest{
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"12345678-1234-1234-123456789012",
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewChatRequestValidation()

			got := target.CreateRoom(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestChatRequestValidation_CreateMessage(t *testing.T) {
	t.Parallel()
	type args struct {
		req *chat.CreateMessageRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &chat.CreateMessageRequest{
					RoomId: "12345678-1234-1234-123456789012",
					UserId: "12345678-1234-1234-123456789012",
					Text:   "テストメッセージ",
				},
			},
			want: true,
		},
		{
			name: "validation error: RoomId.min_len",
			args: args{
				req: &chat.CreateMessageRequest{
					RoomId: "",
					UserId: "12345678-1234-1234-123456789012",
					Text:   "テストメッセージ",
				},
			},
			want: false,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &chat.CreateMessageRequest{
					RoomId: "12345678-1234-1234-123456789012",
					UserId: "",
					Text:   "テストメッセージ",
				},
			},
			want: false,
		},
		{
			name: "validation error: Text.min_len",
			args: args{
				req: &chat.CreateMessageRequest{
					RoomId: "12345678-1234-1234-123456789012",
					UserId: "12345678-1234-1234-123456789012",
					Text:   "",
				},
			},
			want: false,
		},
		{
			name: "validation error: Text.max_len",
			args: args{
				req: &chat.CreateMessageRequest{
					RoomId: "12345678-1234-1234-123456789012",
					UserId: "12345678-1234-1234-123456789012",
					Text:   strings.Repeat("x", 1001),
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewChatRequestValidation()

			got := target.CreateMessage(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestChatRequestValidation_UploadChatImage(t *testing.T) {
	t.Parallel()
	type args struct {
		req *chat.UploadChatImageRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &chat.UploadChatImageRequest{
					UserId:   "12345678-1234-1234-123456789012",
					Image:    []byte{},
					Position: 0,
				},
			},
			want: true,
		},
		{
			name: "validation error: Position.gte",
			args: args{
				req: &chat.UploadChatImageRequest{
					UserId:   "12345678-1234-1234-123456789012",
					Image:    []byte{},
					Position: -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewChatRequestValidation()

			got := target.UploadChatImage(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}
