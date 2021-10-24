package v1

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestAuthResponse(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	type args struct {
		auth *entity.Auth
	}
	tests := []struct {
		name   string
		args   args
		expect *AuthResponse
	}{
		{
			name: "success",
			args: args{
				auth: &entity.Auth{
					Auth: &user.Auth{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						Email:            "test-user@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						Role:             user.Role_ROLE_ADMIN,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
			expect: &AuthResponse{
				ID:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             entity.Role(user.Gender_GENDER_MAN),
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAuthResponse(tt.args.auth)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAuthThumbnailResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		thumbnailURL string
	}
	tests := []struct {
		name   string
		args   args
		expect *AuthThumbnailResponse
	}{
		{
			name: "success",
			args: args{
				thumbnailURL: "https://go.dev/images/gophers/ladder.svg",
			},
			expect: &AuthThumbnailResponse{
				ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAuthThumbnailResponse(tt.args.thumbnailURL)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
