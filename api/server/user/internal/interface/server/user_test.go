package server

import (
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

func testUser(id string) *user.User {
	current := time.Now().Local()

	return &user.User{
		ID:               id,
		Username:         "テストユーザー",
		Gender:           user.MaleGender,
		Email:            "test-user@calmato.jp",
		PhoneNumber:      "000-0000-0000",
		Role:             user.UserRole,
		ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
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
		InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
		CreatedAt:        current,
		UpdatedAt:        current,
	}
}
