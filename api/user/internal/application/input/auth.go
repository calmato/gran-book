package input

// CreateAuth - ユーザ登録のリクエスト
type CreateAuth struct {
	Username             string `validate:"required,max=32"`
	Email                string `validate:"required,email,max=256"`
	Password             string `validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `validate:"required,eqfield=Password"`
}

// UpdateAuthPassword - パスワード変更のリクエスト
type UpdateAuthPassword struct {
	Password             string `validate:"required,password,min=6,max=32"`
	PasswordConfirmation string `validate:"required,eqfield=Password"`
}
