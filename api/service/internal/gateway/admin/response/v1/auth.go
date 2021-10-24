package v1

import "github.com/calmato/gran-book/api/service/internal/gateway/entity"

// 認証情報
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

func NewAuthResponse(a *entity.Auth) *AuthResponse {
	return &AuthResponse{
		ID:               a.Id,
		Username:         a.Username,
		Email:            a.Email,
		PhoneNumber:      a.PhoneNumber,
		Role:             a.Role(),
		ThumbnailURL:     a.ThumbnailUrl,
		SelfIntroduction: a.SelfIntroduction,
		LastName:         a.LastName,
		FirstName:        a.FirstName,
		LastNameKana:     a.LastNameKana,
		FirstNameKana:    a.FirstNameKana,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
	}
}

// サムネイルURL
type AuthThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func NewAuthThumbnailResponse(thumbnailURL string) *AuthThumbnailResponse {
	return &AuthThumbnailResponse{
		ThumbnailURL: thumbnailURL,
	}
}
