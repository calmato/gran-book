package errors

import "errors"

var (
	ErrBadRequest         = errors.New("errors: bad request")           // 400
	ErrUnauthorized       = errors.New("errors: unauthorized")          // 401
	ErrForbidden          = errors.New("errors: forbidden")             // 403
	ErrNotFound           = errors.New("errors: not found")             // 404
	ErrConflict           = errors.New("errors: conflict")              // 409
	ErrPreconditionFailed = errors.New("errors: precondition failed")   // 412
	ErrTooManyRequests    = errors.New("errors: too many requests")     // 429
	ErrInternal           = errors.New("errors: internal server error") // 500
	ErrServiceUnabailable = errors.New("errors: service unavailable")   // 503
	ErrGatewayTimeout     = errors.New("errors: gateway timeout")       // 504
	ErrUnknown            = errors.New("errors: unknown error")         // 500
)
