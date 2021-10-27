package util

import (
	"net/http"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode int                 `json:"status"`  // ステータスコード
	ErrorCode  exception.ErrorCode `json:"code"`    // エラーコード
	Message    string              `json:"message"` // エラー概要
	Detail     string              `json:"detail"`  // エラー詳細
}

func ErrorHandling(ctx *gin.Context, err error) {
	// gRPCコードに変換 (ok: gRPCのレスポンス, ng: その他)
	st, ok := status.FromError(err)
	if ok {
		ec := convertStatusGrpcToHTTP(st)
		err = ec.New(err)
	}

	code, res := getHTTPError(err)
	ctx.JSON(code, res)
	ctx.Abort()
}

func IsNotFound(err error) bool {
	// gRPCコードに変換 (ok: gRPCのレスポンス, ng: その他)
	st, ok := status.FromError(err)
	if !ok {
		return false
	}

	ec := convertStatusGrpcToHTTP(st)
	return ec == exception.ErrNotFound
}

func getHTTPError(err error) (int, *ErrorResponse) {
	res := &ErrorResponse{
		ErrorCode: getErrorCode(err),
		Detail:    getError(err),
	}

	switch getErrorCode(err) {
	case exception.ErrInvalidArgument:
		res.StatusCode = http.StatusBadRequest
		res.Message = "bad request"
	case exception.ErrUnauthorized:
		res.StatusCode = http.StatusUnauthorized
		res.Message = "unauthenticated"
	case exception.ErrForbidden:
		res.StatusCode = http.StatusForbidden
		res.Message = "forbidden"
	case exception.ErrNotFound:
		res.StatusCode = http.StatusNotFound
		res.Message = "not found"
	case exception.ErrConflict:
		res.StatusCode = http.StatusConflict
		res.Message = "conflict"
	case exception.ErrPreconditionFailed:
		res.StatusCode = http.StatusPreconditionFailed
		res.Message = "precondition failed"
	case exception.ErrInternal:
		res.StatusCode = http.StatusInternalServerError
		res.Message = "internal server error"
	case exception.ErrNotImplemented:
		res.StatusCode = http.StatusNotImplemented
		res.Message = "not implemented"
	case exception.ErrServiceUnavailable:
		res.StatusCode = http.StatusServiceUnavailable
		res.Message = "service unavailable"
	case exception.ErrGatewayTimeout:
		res.StatusCode = http.StatusGatewayTimeout
		res.Message = "gateway timeout"
	default:
		res.StatusCode = http.StatusInternalServerError
		res.Message = "unknown"
	}

	return res.StatusCode, res
}

func convertStatusGrpcToHTTP(st *status.Status) exception.ErrorCode {
	switch st.Code() {
	case codes.Unknown:
		return exception.ErrInternal
	case codes.InvalidArgument:
		return exception.ErrInvalidArgument
	case codes.DeadlineExceeded:
		return exception.ErrGatewayTimeout
	case codes.NotFound:
		return exception.ErrNotFound
	case codes.AlreadyExists:
		return exception.ErrConflict
	case codes.PermissionDenied:
		return exception.ErrForbidden
	case codes.FailedPrecondition:
		return exception.ErrPreconditionFailed
	case codes.Unimplemented:
		return exception.ErrNotImplemented
	case codes.Internal:
		return exception.ErrInternal
	case codes.Unavailable:
		return exception.ErrServiceUnavailable
	case codes.Unauthenticated:
		return exception.ErrUnauthorized
	default:
		return exception.ErrInternal
	}
}

func getError(err error) string {
	if e, ok := err.(exception.CustomError); ok {
		return e.Error()
	}

	return ""
}

func getErrorCode(err error) exception.ErrorCode {
	if e, ok := err.(exception.CustomError); ok {
		return e.Code()
	}

	return exception.ErrInternal
}
