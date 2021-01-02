package exception

// バリデーションタグ
const (
	RequiredTag = "required"
	EqFieldTag  = "eqfield"
	MinTag      = "min"
	MaxTag      = "max"
	EmailTag    = "email"
	PasswordTag = "password"
	UniqueTag   = "unique"
)

// バリデーションメッセージ
const (
	RequiredMessage = "is required"
	EqFieldMessage  = "does not match %s"
	MinMessage      = "must be at least %s characters"
	MaxMessage      = "must be at %s or less characters"
	EmailMessage    = "should be in email format"
	PasswordMessage = "should be in password format"
	UniqueMessage   = "must be unique"
)

// カスタムバリデーションメッセージ
const (
	CustomUniqueMessage        = "already exists"
	UnableConvertBase64Massage = "cannot be converted to base64"
)
