package exception

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	ErrorCode        ErrorCode
	Value            error
	ValidationErrors []*ValidationError
}

// ValidationError - バリデーションエラー用構造体
type ValidationError struct {
	Field   string
	Message string
}

// ErrorCode - システムエラーの種類
type ErrorCode uint

const (
	// Unknown - 不明なエラー
	Unknown ErrorCode = iota
	// ErrorInDatastore - データストアでのエラー
	ErrorInDatastore
	// ErrorInStorage - ストレージでのエラー
	ErrorInStorage
	// ErrorInOtherAPI - 他のAPIでのエラー
	ErrorInOtherAPI
	// InvalidDomainValidation - ドメイン層でのバリデーションエラー
	InvalidDomainValidation
	// InvalidRequestValidation - リクエスト値のバリデーションエラ
	InvalidRequestValidation
	// InvalidDatabaseTransaction - データベースでのトランザクションエラー
	InvalidDatabaseTransaction
	// UnableConvertBase64 - Byte64型への変換エラー
	UnableConvertBase64
	// Unauthorized - 認証エラー
	Unauthorized
	// Expired - セッションが有効期限切れ
	Expired
	// Forbidden - 権限エラー
	Forbidden
	// NotFound - 取得エラー
	NotFound
	// NotExistsInDatastore - データストアに対象のレコードが存在しない
	NotExistsInDatastore
	// NotExistsInStorage - ストレージに対象のファイルが存在しない
	NotExistsInStorage
	// Conflict - 既存レコード,ファイルとの競合
	Conflict
)

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error, ves ...*ValidationError) error {
	return CustomError{
		ErrorCode:        ec,
		Value:            err,
		ValidationErrors: ves,
	}
}

// Error - エラー内容を返す
func (ce CustomError) Error() string {
	return ce.Value.Error()
}

// Code - エラーコードを返す
func (ce CustomError) Code() ErrorCode {
	return ce.ErrorCode
}

// Validations - バリデーションエラーの詳細を返す
func (ce CustomError) Validations() []*ValidationError {
	return ce.ValidationErrors
}
