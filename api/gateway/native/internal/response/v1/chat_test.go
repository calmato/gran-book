package v1

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestChatRoomResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		room  *entity.ChatRoom
		users map[string]*entity.User
	}
	tests := []struct {
		name   string
		args   args
		expect *ChatRoomResponse
	}{
		{
			name: "success",
			args: args{
				room: &entity.ChatRoom{
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
							CreatedAt: test.TimeMock,
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
						},
					},
				},
			},
			expect: &ChatRoomResponse{
				ID: "00000000-0000-0000-0000-000000000000",
				Users: []*chatRoomUser{
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
				LatestMessage: &chatRoomMessage{
					UserID:    "12345678-1234-1234-1234-123456789012",
					Text:      "テストメッセージです。",
					Image:     "",
					CreatedAt: test.TimeMock,
				},
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
		{
			name: "success latest message is nil",
			args: args{
				room: &entity.ChatRoom{
					Room: &chat.Room{
						Id: "00000000-0000-0000-0000-000000000000",
						UserIds: []string{
							"12345678-1234-1234-1234-123456789012",
							"23456789-2345-2345-2345-234567890123",
						},
						LatestMessage: nil,
						CreatedAt:     test.TimeMock,
						UpdatedAt:     test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
						},
					},
				},
			},
			expect: &ChatRoomResponse{
				ID: "00000000-0000-0000-0000-000000000000",
				Users: []*chatRoomUser{
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
				LatestMessage: &chatRoomMessage{},
				CreatedAt:     test.TimeMock,
				UpdatedAt:     test.TimeMock,
			},
		},
		{
			name: "success users is length 0",
			args: args{
				room: &entity.ChatRoom{
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
							CreatedAt: test.TimeMock,
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
				users: map[string]*entity.User{},
			},
			expect: &ChatRoomResponse{
				ID: "00000000-0000-0000-0000-000000000000",
				Users: []*chatRoomUser{
					{
						ID:       "12345678-1234-1234-1234-123456789012",
						Username: "unknown",
					},
					{
						ID:       "23456789-2345-2345-2345-234567890123",
						Username: "unknown",
					},
				},
				LatestMessage: &chatRoomMessage{
					UserID:    "12345678-1234-1234-1234-123456789012",
					Text:      "テストメッセージです。",
					Image:     "",
					CreatedAt: test.TimeMock,
				},
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRoomResponse(tt.args.room, tt.args.users)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChatRoomListResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		rooms entity.ChatRooms
		users map[string]*entity.User
	}
	tests := []struct {
		name   string
		args   args
		expect *ChatRoomListResponse
	}{
		{
			name: "success",
			args: args{
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
								CreatedAt: test.TimeMock,
							},
							CreatedAt: test.TimeMock,
							UpdatedAt: test.TimeMock,
						},
					},
					{
						Room: &chat.Room{
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
						},
					},
				},
			},
			expect: &ChatRoomListResponse{
				Rooms: []*chatRoomListRoom{
					{
						ID: "00000000-0000-0000-0000-000000000000",
						Users: []*chatRoomListUser{
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
						LatestMessage: &chatRoomListMessage{
							UserID:    "12345678-1234-1234-1234-123456789012",
							Text:      "テストメッセージです。",
							Image:     "",
							CreatedAt: test.TimeMock,
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
					{
						ID: "11111111-1111-1111-1111-111111111111",
						Users: []*chatRoomListUser{
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
						LatestMessage: &chatRoomListMessage{},
						CreatedAt:     test.TimeMock,
						UpdatedAt:     test.TimeMock,
					},
				},
			},
		},
		{
			name: "success user is length 0",
			args: args{
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
								CreatedAt: test.TimeMock,
							},
							CreatedAt: test.TimeMock,
							UpdatedAt: test.TimeMock,
						},
					},
					{
						Room: &chat.Room{
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
				users: map[string]*entity.User{},
			},
			expect: &ChatRoomListResponse{
				Rooms: []*chatRoomListRoom{
					{
						ID: "00000000-0000-0000-0000-000000000000",
						Users: []*chatRoomListUser{
							{
								ID:       "12345678-1234-1234-1234-123456789012",
								Username: "unknown",
							},
							{
								ID:       "23456789-2345-2345-2345-234567890123",
								Username: "unknown",
							},
						},
						LatestMessage: &chatRoomListMessage{
							UserID:    "12345678-1234-1234-1234-123456789012",
							Text:      "テストメッセージです。",
							Image:     "",
							CreatedAt: test.TimeMock,
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
					{
						ID: "11111111-1111-1111-1111-111111111111",
						Users: []*chatRoomListUser{
							{
								ID:       "12345678-1234-1234-1234-123456789012",
								Username: "unknown",
							},
							{
								ID:       "23456789-2345-2345-2345-234567890123",
								Username: "unknown",
							},
						},
						LatestMessage: &chatRoomListMessage{},
						CreatedAt:     test.TimeMock,
						UpdatedAt:     test.TimeMock,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatRoomListResponse(tt.args.rooms, tt.args.users)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChatMessageResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		message *entity.ChatMessage
		auth    *entity.Auth
	}
	tests := []struct {
		name   string
		args   args
		expect *ChatMessageResponse
	}{
		{
			name: "success",
			args: args{
				message: &entity.ChatMessage{
					Message: &chat.Message{
						Id:        "00000000-0000-0000-0000-000000000000",
						UserId:    "12345678-1234-1234-1234-123456789012",
						Text:      "テストメッセージです。",
						Image:     "",
						CreatedAt: test.TimeMock,
					},
				},
				auth: &entity.Auth{
					Auth: &user.Auth{
						Id:               "12345678-1234-1234-1234-123456789012",
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						Email:            "test-user@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						Role:             user.Role_ROLE_USER,
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						PostalCode:       "000-0000",
						Prefecture:       "東京都",
						City:             "小金井市",
						AddressLine1:     "貫井北町4-1-1",
						AddressLine2:     "",
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
			},
			expect: &ChatMessageResponse{
				Text:  "テストメッセージです。",
				Image: "",
				User: &chatMessageUser{
					ID:           "12345678-1234-1234-1234-123456789012",
					Username:     "テストユーザー",
					ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
				},
				CreatedAt: test.TimeMock,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewChatMessageResponse(tt.args.message, tt.args.auth)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
