package entity

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	code  ErrorCode
	value error
}

// ErrorCode - システムエラーの種類
type ErrorCode int64

const (
	ErrBadRequest          ErrorCode = iota // バリデーションエラー
	ErrUnauthenticated                      // 認証エラー
	ErrForbidden                            // 権限エラー
	ErrNotFound                             // 取得エラー
	ErrConflict                             // ユニークチェックエラー
	ErrInternalServerError                  // サーバーエラー
	ErrNotImplemented                       // 未実装エラー
	ErrServiceUnavailable                   // 到達エラー
)

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error) error {
	return CustomError{
		code:  ec,
		value: err,
	}
}

// Error - エラー内容を返す
func (ce CustomError) Error() string {
	return ce.value.Error()
}

// Code - エラーコードを返す
func (ce CustomError) Code() ErrorCode {
	return ce.code
}
