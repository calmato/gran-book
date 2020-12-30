package v1

import (
	"fmt"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandling(e error) error {
	res := toGRPCError(e)

	return res
}

func toGRPCError(e error) error {
	var c codes.Code

	switch getErrorCode(e) {
	case exception.InvalidRequestValidation, exception.UnableConvertBase64:
		c = codes.InvalidArgument
	case exception.NotFound:
		c = codes.NotFound
	case exception.Conflict:
		c = codes.AlreadyExists
	case exception.Forbidden:
		c = codes.PermissionDenied
	case exception.ErrorInDatastore, exception.ErrorInStorage,
		exception.ErrorInOtherAPI, exception.InvalidDomainValidation:
		c = codes.Internal
	case exception.Unauthorized, exception.Expired:
		c = codes.Unauthenticated
	default:
		c = codes.Unknown
	}

	st := status.New(c, c.String())

	switch c {
	case codes.InvalidArgument: // Bad Requestのとき、追加のエラーメッセージ
		br := &errdetails.BadRequest{
			FieldViolations: getValidationErrors(e),
		}

		ds, err := st.WithDetails(br)
		if err != nil {
			st = status.New(codes.Unknown, fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}

		return ds.Err()
	case codes.Internal: // Internal Server Errorのとき、詳細の追加
		br := &errdetails.LocalizedMessage{
			Locale:  "ja-JP",
			Message: getError(e),
		}

		ds, err := st.WithDetails(br)
		if err != nil {
			st = status.New(codes.Unknown, fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}

		return ds.Err()
	default:
		return st.Err()
	}
}

func getError(e error) string {
	if err, ok := e.(exception.CustomError); ok {
		return err.Error()
	}

	return ""
}

func getErrorCode(e error) exception.ErrorCode {
	if err, ok := e.(exception.CustomError); ok {
		return err.Code()
	}

	return exception.Unknown
}

func getValidationErrors(e error) []*errdetails.BadRequest_FieldViolation {
	if err, ok := e.(exception.CustomError); ok {
		ves := make([]*errdetails.BadRequest_FieldViolation, len(err.Validations()))
		for i, ve := range err.Validations() {
			ves[i] = &errdetails.BadRequest_FieldViolation{
				Field:       ve.Field,
				Description: ve.Message,
			}
		}

		return ves
	}

	return []*errdetails.BadRequest_FieldViolation{}
}
