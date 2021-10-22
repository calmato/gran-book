package server

import (
	"errors"

	"github.com/calmato/gran-book/api/service/pkg/exception"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errInvalidUploadRequest = errors.New("server: position is duplicated")

func toGRPCError(e error) error {
	c := getGRPCErrorCode(getErrorCode(e))

	st := status.New(c, c.String())

	switch c {
	case codes.InvalidArgument, codes.AlreadyExists:
		br := &errdetails.BadRequest{
			FieldViolations: getValidationErrors(e),
		}

		dt, _ := st.WithDetails(br)
		return dt.Err()
	case codes.Internal, codes.Unknown:
		br := &errdetails.LocalizedMessage{
			Locale:  "en-US",
			Message: getError(e),
		}

		dt, _ := st.WithDetails(br)
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

	return exception.ErrUnknown
}

func getGRPCErrorCode(code exception.ErrorCode) codes.Code {
	switch code {
	case exception.ErrInvalidArgument,
		exception.ErrInvalidRequestValidation,
		exception.ErrUnableConvertBase64:
		return codes.InvalidArgument
	case exception.ErrGatewayTimeout:
		return codes.DeadlineExceeded
	case exception.ErrNotExistsInDatastore,
		exception.ErrNotExistsInStorage,
		exception.ErrNotFound:
		return codes.NotFound
	case exception.ErrConflict:
		return codes.AlreadyExists
	case exception.ErrForbidden:
		return codes.PermissionDenied
	case exception.ErrPreconditionFailed:
		return codes.FailedPrecondition
	case exception.ErrNotImplemented:
		return codes.Unimplemented
	case exception.ErrInDatastore,
		exception.ErrInStorage,
		exception.ErrInvalidDomainValidation:
		return codes.Internal
	case exception.ErrSessionExpired,
		exception.ErrUnauthorized:
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
