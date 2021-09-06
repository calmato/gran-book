package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		auth         *Auth
		expectGender Gender
	}{
		{
			name: "success",
			auth: &Auth{
				Auth: &pb.Auth{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           pb.Gender_GENDER_MAN,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             pb.Role_ROLE_USER,
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
			expectGender: GenderMan,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectGender, tt.auth.Gender())
		})
	}
}

func TestUsers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		users     Users
		expectMap map[string]*User
	}{
		{
			name: "success",
			users: Users{
				{
					User: &pb.User{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           pb.Gender_GENDER_MAN,
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
				{
					User: &pb.User{
						Id:               "11111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー",
						Gender:           pb.Gender_GENDER_MAN,
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
			expectMap: map[string]*User{
				"00000000-0000-0000-0000-000000000000": {
					User: &pb.User{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           pb.Gender_GENDER_MAN,
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
				"11111111-1111-1111-1111-111111111111": {
					User: &pb.User{
						Id:               "11111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー",
						Gender:           pb.Gender_GENDER_MAN,
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectMap, tt.users.Map())
		})
	}
}
