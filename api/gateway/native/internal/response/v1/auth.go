package v1

type AuthResponse struct {
	ID               string `json:"id"`               // ユーザーID
	Username         string `json:"username"`         // ユーザー名
	Gender           int    `json:"gender"`           // 性別
	Email            string `json:"email"`            // メールアドレス
	PhoneNumber      string `json:"phoneNumber"`      // 電話番号
	Role             int    `json:"role"`             // ユーザー権限
	ThumbnailURL     string `json:"thumbnailUrl"`     // サムネイルURL
	SelfIntroduction string `json:"selfIntroduction"` // 自己紹介
	LastName         string `json:"lastName"`         // 姓
	FirstName        string `json:"firstName"`        // 名
	LastNameKana     string `json:"lastNameKana"`     // 姓(かな)
	FirstNameKana    string `json:"firstNameKana"`    // 名(かな)
	PostalCode       string `json:"postalCode"`       // 郵便番号
	Prefecture       string `json:"prefecture"`       // 都道府県
	City             string `json:"city"`             // 市区町村
	AddressLine1     string `json:"addressLine1"`     // 町名,番地
	AddressLine2     string `json:"addressLine2"`     // マンション・ビル名,号室
	CreatedAt        string `json:"createdAt"`        // 作成日時
	UpdatedAt        string `json:"updatedAt"`        // 更新日時
}

type AuthThumbnailResponse struct {
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
