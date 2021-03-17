package input

// CreateAuth - ユーザ登録のリクエスト
type CreateAuth struct {
	Username             string `json:"username" validate:"required,max=32"`
	Email                string `json:"email" validate:"required,email,max=256"`
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// UpdateAuthEmail - メールアドレス変更のリクエスト
type UpdateAuthEmail struct {
	Email string `json:"email" validate:"required,email,max=256"`
}

// UpdateAuthPassword - パスワード変更のリクエスト
type UpdateAuthPassword struct {
	Password             string `json:"password" validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// UpdateAuthProfile - プロフィール変更のリクエスト
type UpdateAuthProfile struct {
	Username         string `json:"username" validate:"required,max=32"`
	Gender           int    `json:"gender" validate:"gte=0,lte=2"`
	Thumbnail        string `json:"thumbnail"`
	SelfIntroduction string `json:"selfIntroduction" validate:"max=256"`
}

// UpdateAuthAddress - 住所変更のリクエスト
type UpdateAuthAddress struct {
	LastName      string `json:"lastName" validate:"required,max=16"`
	FirstName     string `json:"firstName" validate:"required,max=16"`
	LastNameKana  string `json:"lastNameKana" validate:"required,hiragana,max=32"`
	FirstNameKana string `json:"firstNameKana" validate:"required,hiragana,max=32"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,max=16"`
	PostalCode    string `json:"postalCode" validate:"required,max=16"`
	Prefecture    string `json:"prefecture" validate:"required,max=32"`
	City          string `json:"city" validate:"required,max=32"`
	AddressLine1  string `json:"addressLine1" validate:"required,max=64"`
	AddressLine2  string `json:"addressLine2" validate:"max=64"`
}
