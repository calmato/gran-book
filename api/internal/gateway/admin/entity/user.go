package entity

import "github.com/calmato/gran-book/api/internal/gateway/entity"

type User struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // ユーザー名
	Email            string `json:"email"`            // メールアドレス
	PhoneNumber      string `json:"phoneNumber"`      // 電話番号
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	LastName         string `json:"lastName"`         // 姓
	FirstName        string `json:"firstName"`        // 名
	LastNameKana     string `json:"lastNameKana"`     // 姓(かな)
	FirstNameKana    string `json:"firstNameKana"`    // 名(かな)
	CreatedAt        string `json:"createdAt"`        // 作成日時
	UpdatedAt        string `json:"updatedAt"`        // 更新日時
}

type Users []*User

func NewUser(u *entity.User) *User {
	return &User{
		ID:               u.Id,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		ThumbnailURL:     u.ThumbnailUrl,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

func NewUsers(us entity.Users) Users {
	res := make(Users, len(us))
	for i := range us {
		res[i] = NewUser(us[i])
	}
	return res
}
