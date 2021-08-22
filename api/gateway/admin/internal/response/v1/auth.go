package v1

import "github.com/calmato/gran-book/api/gateway/admin/internal/entity"

type AuthResponse struct {
	ID               string      `json:"id"`               // ユーザーID
	Username         string      `json:"username"`         // ユーザー名
	Email            string      `json:"email"`            // メールアドレス
	PhoneNumber      string      `json:"phoneNumber"`      // 電話番号
	Role             entity.Role `json:"role"`             // ユーザー権限
	ThumbnailURL     string      `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string      `json:"selfIntroduction"` // 自己紹介
	LastName         string      `json:"lastName"`         // 姓
	FirstName        string      `json:"firstName"`        // 名
	LastNameKana     string      `json:"lastNameKana"`     // 姓(かな)
	FirstNameKana    string      `json:"firstNameKana"`    // 名(かな)
	CreatedAt        string      `json:"createdAt"`        // 作成日時
	UpdatedAt        string      `json:"updatedAt"`        // 更新日時
}

type AuthThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
