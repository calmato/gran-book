package v1

import "github.com/calmato/gran-book/api/internal/gateway/entity"

// 新規登録
type CreateAuthRequest struct {
	Username             string `json:"username" binding:"required"`             // 表示名
	Email                string `json:"email" binding:"required"`                // メールアドレス
	Password             string `json:"password" binding:"required"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"` // パスワード(確認用)
}

// メールアドレス更新
type UpdateAuthEmailRequest struct {
	Email string `json:"email" binding:"required"` // メールアドレス
}

// パスワード更新
type UpdateAuthPasswordRequest struct {
	Password             string `json:"password" binding:"required"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"` // パスワード(確認用)
}

// プロフィール更新
type UpdateAuthProfileRequest struct {
	Username         string        `json:"username" binding:"required"` // 表示名
	Gender           entity.Gender `json:"gender" binding:""`           // 性別
	ThumbnailURL     string        `json:"thumbnailUrl" binding:""`     // サムネイルURL
	SelfIntroduction string        `json:"selfIntroduction" binding:""` // 自己紹介
}

// 住所更新
type UpdateAuthAddressRequest struct {
	LastName      string `json:"lastName" binding:"required"`      // 姓
	FirstName     string `json:"firstName" binding:"required"`     // 名
	LastNameKana  string `json:"lastNameKana" binding:"required"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana" binding:"required"` // 名(かな)
	PhoneNumber   string `json:"phoneNumber" binding:"required"`   // 電話番号
	PostalCode    string `json:"postalCode" binding:"required"`    // 郵便番号
	Prefecture    string `json:"prefecture" binding:"required"`    // 都道府県
	City          string `json:"city" binding:"required"`          // 市区町村
	AddressLine1  string `json:"addressLine1" binding:"required"`  // 町名,番地
	AddressLine2  string `json:"addressLine2" binding:""`          // マンション・ビル名,号室
}

// 端末ID登録
type RegisterAuthDeviceRequest struct {
	InstanceID string `json:"instanceId" binding:"required"` // 端末ID
}
