package v1

import "github.com/calmato/gran-book/api/gateway/native/internal/entity"

// 認証情報
type AuthV1Response struct {
	Id               string        `json:"id,omitempty"`               // ユーザーID
	Username         string        `json:"username,omitempty"`         // ユーザー名
	Gender           entity.Gender `json:"gender,omitempty"`           // 性別
	Email            string        `json:"email,omitempty"`            // メールアドレス
	PhoneNumber      string        `json:"phoneNumber,omitempty"`      // 電話番号
	ThumbnailUrl     string        `json:"thumbnailUrl,omitempty"`     // サムネイルURL
	SelfIntroduction string        `json:"selfIntroduction,omitempty"` // 自己紹介
	LastName         string        `json:"lastName,omitempty"`         // 姓
	FirstName        string        `json:"firstName,omitempty"`        // 名
	LastNameKana     string        `json:"lastNameKana,omitempty"`     // 姓(かな)
	FirstNameKana    string        `json:"firstNameKana,omitempty"`    // 名(かな)
	PostalCode       string        `json:"postalCode,omitempty"`       // 郵便番号
	Prefecture       string        `json:"prefecture,omitempty"`       // 都道府県
	City             string        `json:"city,omitempty"`             // 市区町村
	AddressLine1     string        `json:"addressLine1,omitempty"`     // 町名,番地
	AddressLine2     string        `json:"addressLine2,omitempty"`     // マンション・ビル名,号室
	CreatedAt        string        `json:"createdAt,omitempty"`        // 作成日時
	UpdatedAt        string        `json:"updatedAt,omitempty"`        // 更新日時
}

// サムネイルURL
type AuthThumbnailV1Response struct {
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}
