package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

// 新規登録
type CreateAuthRequest struct {
	Username             string `json:"username"`             // 表示名
	Email                string `json:"email"`                // メールアドレス
	Password             string `json:"password"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation"` // パスワード(確認用)
}

// メールアドレス更新
type UpdateAuthEmailRequest struct {
	Email string `json:"email"` // メールアドレス
}

// パスワード更新
type UpdateAuthPasswordRequest struct {
	Password             string `json:"password"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation"` // パスワード(確認用)
}

// プロフィール更新
type UpdateAuthProfileRequest struct {
	Username         string        `json:"username"`         // 表示名
	Gender           entity.Gender `json:"gender"`           // 性別
	ThumbnailURL     string        `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string        `json:"selfIntroduction"` // 自己紹介
}

// 住所更新
type UpdateAuthAddressRequest struct {
	LastName      string `json:"lastName"`      // 姓
	FirstName     string `json:"firstName"`     // 名
	LastNameKana  string `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana"` // 名(かな)
	PhoneNumber   string `json:"phoneNumber"`   // 電話番号
	PostalCode    string `json:"postalCode"`    // 郵便番号
	Prefecture    string `json:"prefecture"`    // 都道府県
	City          string `json:"city"`          // 市区町村
	AddressLine1  string `json:"addressLine1"`  // 町名,番地
	AddressLine2  string `json:"addressLine2"`  // マンション・ビル名,号室
}

// 端末ID登録
type RegisterAuthDeviceRequest struct {
	InstanceID string `json:"instanceId"` // 端末ID
}
