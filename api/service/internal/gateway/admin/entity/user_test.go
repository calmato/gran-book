package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name   string
		users  entity.Users
		expect Users
	}{
		{
			name: "success",
			users: entity.Users{
				{
					User: &user.User{
						Id:               "00000000-0000-0000-0000-000000000000",
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
			},
			expect: Users{
				{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUsers(tt.users)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
