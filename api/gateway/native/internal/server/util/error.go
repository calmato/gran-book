package util

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	return ec == entity.ErrNotFound
}

func getHTTPError(err error) (int, *response.ErrorResponse) {
	res := &response.ErrorResponse{
		Code:   int64(getErrorCode(err)),
		Detail: getError(err),
	}

	switch getErrorCode(err) {
	case entity.ErrBadRequest:
		res.Status = 400
		res.Message = "bad request"
	case entity.ErrUnauthenticated:
		res.Status = 401
		res.Message = "unauthenticated"
	case entity.ErrForbidden:
		res.Status = 403
		res.Message = "forbidden"
	case entity.ErrNotFound:
		res.Status = 404
		res.Message = "not found"
	case entity.ErrConflict:
		res.Status = 409
		res.Message = "conflict"
	case entity.ErrInternalServerError:
		res.Status = 500
		res.Message = "internal server error"
	case entity.ErrNotImplemented:
		res.Status = 501
		res.Message = "not implemented"
	case entity.ErrServiceUnavailable:
		res.Status = 503
		res.Message = "service unavailable"
	default:
		res.Status = 500
		res.Message = "unknown"
	}

	return int(res.Status), res
}

func convertStatusGrpcToHTTP(st *status.Status) entity.ErrorCode {
	switch st.Code() {
	case codes.InvalidArgument:
		return entity.ErrBadRequest
	case codes.Unauthenticated:
		return entity.ErrUnauthenticated
	case codes.PermissionDenied:
		return entity.ErrForbidden
	case codes.NotFound:
		return entity.ErrNotFound
	case codes.AlreadyExists:
		return entity.ErrConflict
	case codes.Unknown, codes.Internal:
		return entity.ErrInternalServerError
	case codes.Unimplemented:
		return entity.ErrNotImplemented
	case codes.Unavailable:
		return entity.ErrServiceUnavailable
	default:
		return entity.ErrInternalServerError
	}
}

func getError(err error) string {
	if e, ok := err.(entity.CustomError); ok {
		return e.Error()
	}

	return ""
}

func getErrorCode(err error) entity.ErrorCode {
	if e, ok := err.(entity.CustomError); ok {
		return e.Code()
	}

	return entity.ErrInternalServerError
}
