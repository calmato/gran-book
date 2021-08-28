package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

// 新規登録
type CreateAuthV1Request struct {
	Username             string `json:"username,omitempty"`             // 表示名
	Email                string `json:"email,omitempty"`                // メールアドレス
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}

// メールアドレス更新
type UpdateAuthEmailV1Request struct {
	Email string `json:"email,omitempty"` // メールアドレス
}

// パスワード更新
type UpdateAuthPasswordV1Request struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}

// プロフィール更新
type UpdateAuthProfileV1Request struct {
	Username         string        `json:"username,omitempty"`         // 表示名
	Gender           entity.Gender `json:"gender,omitempty"`           // 性別
	ThumbnailUrl     string        `json:"thumbnailUrl,omitempty"`     // サムネイルURL
	SelfIntroduction string        `json:"selfIntroduction,omitempty"` // 自己紹介
}

// 住所更新
type UpdateAuthAddressV1Request struct {
	LastName      string `json:"lastName,omitempty"`      // 姓
	FirstName     string `json:"firstName,omitempty"`     // 名
	LastNameKana  string `json:"lastNameKana,omitempty"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana,omitempty"` // 名(かな)
	PhoneNumber   string `json:"phoneNumber,omitempty"`   // 電話番号
	PostalCode    string `json:"postalCode,omitempty"`    // 郵便番号
	Prefecture    string `json:"prefecture,omitempty"`    // 都道府県
	City          string `json:"city,omitempty"`          // 市区町村
	AddressLine1  string `json:"addressLine1,omitempty"`  // 町名,番地
	AddressLine2  string `json:"addressLine2,omitempty"`  // マンション・ビル名,号室
}

// 端末ID登録
type RegisterAuthDeviceV1Request struct {
	InstanceId string `json:"instanceId,omitempty"` // 端末ID
}
