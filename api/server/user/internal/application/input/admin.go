package input

// ListAdmin - 管理者一覧のリクエスト
type ListAdmin struct {
	Limit  int64 `json:"limit" validate:"gte=0,lte=1000"`
	Offset int64 `json:"offset" validate:"gte=0"`
}

// CreateAdmin - 管理者登録のリクエスト
type CreateAdmin struct {
	Username             string `json:"username" validate:"required,max=32"`
	Email                string `json:"email" validate:"required,email,max=256"`
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
	Role                 int32  `json:"role" validate:"gte=1,lte=3"`
	LastName             string `json:"lastName" validate:"required,max=16"`
	FirstName            string `json:"firstName" validate:"required,max=16"`
	LastNameKana         string `json:"lastNameKana" validate:"required,hiragana,max=32"`
	FirstNameKana        string `json:"firstNameKana" validate:"required,hiragana,max=32"`
}

// UpdateAdminRole - 管理者権限変更のリクエスト
type UpdateAdminRole struct {
	Role int32 `json:"role" validate:"gte=1,lte=3"`
}

// UpdateAdminPassword - 管理者パスワード変更のリクエスト
type UpdateAdminPassword struct {
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// UpdateAdminProfile - 管理者プロフィール変更のリクエスト
type UpdateAdminProfile struct {
	Username      string `json:"username" validate:"required,max=32"`
	Email         string `json:"email" validate:"required,email,max=256"`
	LastName      string `json:"lastName" validate:"required,max=16"`
	FirstName     string `json:"firstName" validate:"required,max=16"`
	LastNameKana  string `json:"lastNameKana" validate:"required,hiragana,max=32"`
	FirstNameKana string `json:"firstNameKana" validate:"required,hiragana,max=32"`
}
