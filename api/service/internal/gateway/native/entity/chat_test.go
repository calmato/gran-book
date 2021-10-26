package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/chat"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestChatRooms(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name   string
		rooms  entity.ChatRooms
		users  map[string]*entity.User
		expect ChatRooms
	}{
		{
			name: "success",
			rooms: entity.ChatRooms{
				{
					Room: &chat.Room{
						Id: "00000000-0000-0000-0000-000000000000",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						LatestMessage: &chat.Message{
							Id:        "00000000-0000-0000-0000-000000000000",
							UserId:    "12345678-1234-1234-1234-123456789012",
							Text:      "テストメッセージです。",
							Image:     "",
							CreatedAt: now,
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
			users: map[string]*entity.User{
				"12345678-1234-1234-1234-123456789012": {
					User: &user.User{
						Id:               "12345678-1234-1234-1234-123456789012",
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						Email:            "test-user01@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
				"23456789-2345-2345-2345-234567890123": {
					User: &user.User{
						Id:               "23456789-2345-2345-2345-234567890123",
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						Email:            "test-user02@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
			expect: ChatRooms{
				{
					ID: "00000000-0000-0000-0000-000000000000",
					Users: ChatUsers{
						{
							ID:           "12345678-1234-1234-1234-123456789012",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						{
							ID:           "23456789-2345-2345-2345-234567890123",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
					},
					LatestMessage: &ChatMessage{
						Text:   "テストメッセージです。",
						Image:  "",
						UserID: "12345678-1234-1234-1234-123456789012",
						User: &ChatUser{
							ID:           "12345678-1234-1234-1234-123456789012",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						CreatedAt: now,
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID: "11111111-1111-1111-1111-111111111111",
					Users: ChatUsers{
						{
							ID:           "12345678-1234-1234-1234-123456789012",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						{
							ID:           "23456789-2345-2345-2345-234567890123",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
					},
					LatestMessage: &ChatMessage{},
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
		},
		{
			name: "success user is length 0",
			rooms: entity.ChatRooms{
				{
					Room: &chat.Room{
						Id: "00000000-0000-0000-0000-000000000000",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						LatestMessage: &chat.Message{
							Id:        "00000000-0000-0000-0000-000000000000",
							UserId:    "12345678-1234-1234-1234-123456789012",
							Text:      "テストメッセージです。",
							Image:     "",
							CreatedAt: now,
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
			users: map[string]*entity.User{},
			expect: ChatRooms{
				{
					ID: "00000000-0000-0000-0000-000000000000",
					Users: ChatUsers{
						{
							ID:           "",
							Username:     "unknown",
							ThumbnailURL: "",
						},
						{
							ID:           "",
							Username:     "unknown",
							ThumbnailURL: "",
						},
					},
					LatestMessage: &ChatMessage{
						Text:   "テストメッセージです。",
						Image:  "",
						UserID: "12345678-1234-1234-1234-123456789012",
						User: &ChatUser{
							ID:           "",
							Username:     "unknown",
							ThumbnailURL: "",
						},
						CreatedAt: now,
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID: "11111111-1111-1111-1111-111111111111",
					Users: ChatUsers{
						{
							ID:           "",
							Username:     "unknown",
							ThumbnailURL: "",
						},
						{
							ID:           "",
							Username:     "unknown",
							ThumbnailURL: "",
						},
					},
					LatestMessage: &ChatMessage{},
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRooms(tt.rooms, tt.users)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChatMessage(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name    string
		message *entity.ChatMessage
		user    *entity.User
		expect  *ChatMessage
	}{
		{
			name: "success",
			message: &entity.ChatMessage{
				Message: &chat.Message{
					Id:        "00000000-0000-0000-0000-000000000000",
					UserId:    "12345678-1234-1234-1234-123456789012",
					Text:      "テストメッセージです。",
					Image:     "",
					CreatedAt: now,
				},
			},
			user: &entity.User{
				User: &user.User{
					Id:               "12345678-1234-1234-1234-123456789012",
					Username:         "テストユーザー",
					Gender:           user.Gender_GENDER_MAN,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expect: &ChatMessage{
				Text:   "テストメッセージです。",
				Image:  "",
				UserID: "12345678-1234-1234-1234-123456789012",
				User: &ChatUser{
					ID:           "12345678-1234-1234-1234-123456789012",
					Username:     "テストユーザー",
					ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
				},
				CreatedAt: now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatMessage(tt.message, tt.user)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
