package input

// ListAdmin - 管理者一覧のリクエスト
type ListAdmin struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}

// SearchAdmin - 管理者検索のリクエスト
type SearchAdmin struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
	Field     string `json:"field" validate:"required,oneof=username email"`
	Value     string `json:"value" validate:"required"`
}

// CreateAdmin - 管理者登録のリクエスト
type CreateAdmin struct {
	Username             string `json:"username" validate:"required,max=32"`
	Email                string `json:"email" validate:"required,email,max=256"`
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
	Role                 int    `json:"role" validate:"gte=1,lte=3"`
	LastName             string `json:"lastName" validate:"required,max=16"`
	FirstName            string `json:"firstName" validate:"required,max=16"`
	LastNameKana         string `json:"lastNameKana" validate:"required,hiragana,max=32"`
	FirstNameKana        string `json:"firstNameKana" validate:"required,hiragana,max=32"`
}

// UpdateAdminContact - 管理者連絡先変更のリクエスト
type UpdateAdminContact struct {
	Email       string `json:"email" validate:"required,email,max=256"`
	PhoneNumber string `json:"phoneNumber" validate:"required,max=16"`
}

// UpdateAdminPassword - 管理者パスワード変更のリクエスト
type UpdateAdminPassword struct {
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// UpdateAdminProfile - 管理者プロフィール変更のリクエスト
type UpdateAdminProfile struct {
	Username      string `json:"username" validate:"required,max=32"`
	LastName      string `json:"lastName" validate:"required,max=16"`
	FirstName     string `json:"firstName" validate:"required,max=16"`
	LastNameKana  string `json:"lastNameKana" validate:"required,hiragana,max=32"`
	FirstNameKana string `json:"firstNameKana" validate:"required,hiragana,max=32"`
	Role          int    `json:"role" validate:"gte=1,lte=3"`
	ThumbnailURL  string `json:"thumbnailUrl"`
}

// UploadAdminThumbnail - サムネイルアップロードのリクエスト
type UploadAdminThumbnail struct {
	Thumbnail []byte `json:"thumbnail" validate:"required"`
}
