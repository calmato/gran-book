package v1

import (
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
)

func testUser(id string) *user.User {
	return &user.User{
		Id:               id,
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
		CreatedAt:        test.TimeMock,
		UpdatedAt:        test.TimeMock,
	}
}

func testFollow(id string) *user.Follow {
	return &user.Follow{
		Id:               id,
		Username:         "テストユーザー",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollow:         true,
	}
}

func testFollower(id string) *user.Follower {
	return &user.Follower{
		Id:               id,
		Username:         "テストユーザー",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollow:         false,
	}
}
