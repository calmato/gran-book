package v1

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandling(e error) error {
	res := toGRPCError(e)

	return res
}

func toGRPCError(e error) error {
	c := getGRPCErrorCode(e)

	st := status.New(c, c.String())

	switch c {
	case codes.InvalidArgument, codes.AlreadyExists:
		br := &errdetails.BadRequest{
			FieldViolations: getValidationErrors(e),
		}

		dt, err := st.WithDetails(br)
		if err != nil {
			dt = status.New(codes.Unknown, fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}

		return dt.Err()
	case codes.Internal:
		br := &errdetails.LocalizedMessage{
			Locale:  "en-US",
			Message: getError(e),
		}

		dt, err := st.WithDetails(br)
		if err != nil {
			dt = status.New(codes.Unknown, fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}

		return dt.Err()
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

func getGRPCErrorCode(e error) codes.Code {
	switch getErrorCode(e) {
	case exception.InvalidRequestValidation, exception.UnableConvertBase64:
		return codes.InvalidArgument
	case exception.NotFound:
		return codes.NotFound
	case exception.Conflict:
		return codes.AlreadyExists
	case exception.Forbidden:
		return codes.PermissionDenied
	case exception.ErrorInDatastore, exception.ErrorInStorage,
		exception.ErrorInOtherAPI, exception.InvalidDomainValidation:
		return codes.Internal
	case exception.Unauthorized, exception.Expired:
		return codes.Unauthenticated
	default:
		return codes.Unknown
	}
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
