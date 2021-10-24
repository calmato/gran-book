package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestAdmins(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name   string
		admins entity.Admins
		expect Admins
	}{
		{
			name: "success",
			admins: entity.Admins{
				{
					Admin: &user.Admin{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Email:            "test-user@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						Role:             user.Role_ROLE_ADMIN,
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
			expect: Admins{
				{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             entity.RoleAdmin,
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
			actual := NewAdmins(tt.admins)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
