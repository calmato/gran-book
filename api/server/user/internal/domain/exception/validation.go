package exception

// バリデーションタグ
const (
	RequiredTag         = "required"
	EqFieldTag          = "eqfield"
	MinTag              = "min"
	MaxTag              = "max"
	Base64Tag           = "base64"
	EmailTag            = "email"
	HiraganaTag         = "hiragana"
	PasswordTag         = "password"
	UniqueTag           = "unique"
	LessThanTag         = "lt"
	LessThanEqualTag    = "lte"
	GreaterThanTag      = "gt"
	GreaterThanEqualTag = "gte"
	OneOfTag            = "oneof"
)

// バリデーションメッセージ
const (
	RequiredMessage         = "is required"
	EqFieldMessage          = "does not match %s"
	MinMessage              = "must be at least %s characters"
	MaxMessage              = "must be at %s or less characters"
	Base64Message           = "should be in base64 format"
	EmailMessage            = "should be in email format"
	HiraganaMessage         = "should be in hiragana format"
	PasswordMessage         = "should be in password format"
	UniqueMessage           = "must be unique"
	LessThanMessage         = "must be at less than %s"
	LessThanEuqalMessage    = "must be at less than or equal to %s"
	GreaterThanMessage      = "must be at greater than %s"
	GreaterThanEqualMessage = "must be at greater than or equal to %s"
	OneOfMessage            = "must be one of %s"
)

// カスタムバリデーションメッセージ
const (
	CustomUniqueMessage        = "already exists"
	UnableConvertBase64Massage = "cannot be converted to base64"
)
