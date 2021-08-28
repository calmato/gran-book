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
	LastName         string        `json:"last_name,omitempty"`        // 姓
	FirstName        string        `json:"first_name,omitempty"`       // 名
	LastNameKana     string        `json:"last_name_kana,omitempty"`   // 姓(かな)
	FirstNameKana    string        `json:"first_name_kana,omitempty"`  // 名(かな)
	PostalCode       string        `json:"postal_code,omitempty"`      // 郵便番号
	Prefecture       string        `json:"prefecture,omitempty"`       // 都道府県
	City             string        `json:"city,omitempty"`             // 市区町村
	AddressLine1     string        `json:"address_line1,omitempty"`    // 町名,番地
	AddressLine2     string        `json:"address_line2,omitempty"`    // マンション・ビル名,号室
	CreatedAt        string        `json:"created_at,omitempty"`       // 作成日時
	UpdatedAt        string        `json:"updated_at,omitempty"`       // 更新日時
}
