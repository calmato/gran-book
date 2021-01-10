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
