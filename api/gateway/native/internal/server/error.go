package server

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandling(ctx *gin.Context, err error) {
	// gRPCコードに変換 (ok: gRPCのレスポンス, ng: その他)
	st, ok := status.FromError(err)
	if ok {
		ec := convertStatusGrpcToHttp(st)
		err = ec.New(err)
	}

	code, res := getHTTPError(err)
	ctx.JSON(code, res)
	ctx.Abort()
}

func getHTTPError(err error) (int, *response.ErrorResponse) {
	res := &response.ErrorResponse{}
	res.ErrorCode = getErrorCode(err)
	res.Detail = getError(err)

	switch res.ErrorCode {
	case entity.ErrBadRequest:
		res.StatusCode = 400
		res.Message = "bad request"
	case entity.ErrUnauthenticated:
		res.StatusCode = 401
		res.Message = "unauthenticated"
	case entity.ErrForbidden:
		res.StatusCode = 403
		res.Message = "forbidden"
	case entity.ErrNotFound:
		res.StatusCode = 404
		res.Message = "not found"
	case entity.ErrConflict:
		res.StatusCode = 409
		res.Message = "conflict"
	case entity.ErrInternalServerError:
		res.StatusCode = 500
		res.Message = "internal server error"
	case entity.ErrNotImplemented:
		res.StatusCode = 501
		res.Message = "not implemented"
	case entity.ErrServiceUnavailable:
		res.StatusCode = 503
		res.Message = "service unavailable"
	default:
		res.StatusCode = 500
		res.Message = "unknown"
	}

	return res.StatusCode, res
}

func convertStatusGrpcToHttp(st *status.Status) entity.ErrorCode {
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
