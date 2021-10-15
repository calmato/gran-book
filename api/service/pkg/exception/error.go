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
	// ErrUnknown - 不明なエラー
	ErrUnknown ErrorCode = iota
	// ErrInvalidArgument - バリデーションエラー
	ErrInvalidArgument
	// ErrInvalidRequestValidation - バリデーションエラー (リクエスト)
	ErrInvalidRequestValidation
	// ErrInvalidDomainValidation - バリデーションエラー (ドメイン)
	ErrInvalidDomainValidation
	// ErrUnableConvertBase64 - Byte64型への変換エラー
	ErrUnableConvertBase64
	// ErrUnauthorized - 認証エラー
	ErrUnauthorized
	// ErrSessionExpired - セッションが有効期限切れ
	ErrSessionExpired
	// ErrForbidden - 権限エラー
	ErrForbidden
	// ErrPreconditionFailed - アクセス拒否
	ErrPreconditionFailed
	// ErrNotFound - 取得エラー
	ErrNotFound
	// ErrNotExistsInDatastore - データストアに対象のレコードが存在しない
	ErrNotExistsInDatastore
	// ErrNotExistsInStorage - ストレージに対象のファイルが存在しない
	ErrNotExistsInStorage
	// ErrConflict - 既存レコード,ファイルとの競合
	ErrConflict
	// ErrNotImplemented - サービスが利用不可
	ErrNotImplemented
	// ErrGatewayTimeout - 接続タイムアウト
	ErrGatewayTimeout
	// ErrInternal - 内部エラー
	ErrInternal
	// ErrInDatastore - データベースエラー
	ErrInDatastore
	// ErrInStorage - ストレージエラー
	ErrInStorage
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
