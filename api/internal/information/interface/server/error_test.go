package server

import (
	"testing"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestToGRPCError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    exception.CustomError
		expect func() error
	}{
		{
			name: "success from InvalidArgument",
			err: exception.CustomError{
				ErrorCode: exception.ErrInvalidRequestValidation,
				Value:     test.ErrMock,
				ValidationErrors: []*exception.ValidationError{
					{
						Field:   "id",
						Message: exception.RequiredMessage,
					},
				},
			},
			expect: func() error {
				br := &errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{
							Field:       "id",
							Description: exception.RequiredMessage,
						},
					},
				}
				st := status.New(codes.InvalidArgument, codes.InvalidArgument.String())
				dt, _ := st.WithDetails(br)
				return dt.Err()
			},
		},
		{
			name: "success from DeadlineExceeded",
			err: exception.CustomError{
				ErrorCode: exception.ErrGatewayTimeout,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.DeadlineExceeded, codes.DeadlineExceeded.String()).Err()
			},
		},
		{
			name: "success from NotFound",
			err: exception.CustomError{
				ErrorCode: exception.ErrNotFound,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.NotFound, codes.NotFound.String()).Err()
			},
		},
		{
			name: "success from AlreadyExists",
			err: exception.CustomError{
				ErrorCode: exception.ErrConflict,
				Value:     test.ErrMock,
			},
			expect: func() error {
				br := &errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{},
				}
				st := status.New(codes.AlreadyExists, codes.AlreadyExists.String())
				dt, _ := st.WithDetails(br)
				return dt.Err()
			},
		},
		{
			name: "success from PermissionDenied",
			err: exception.CustomError{
				ErrorCode: exception.ErrForbidden,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.PermissionDenied, codes.PermissionDenied.String()).Err()
			},
		},
		{
			name: "success from FailedPrecondition",
			err: exception.CustomError{
				ErrorCode: exception.ErrPreconditionFailed,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.FailedPrecondition, codes.FailedPrecondition.String()).Err()
			},
		},
		{
			name: "success from Unimplemented",
			err: exception.CustomError{
				ErrorCode: exception.ErrNotImplemented,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
			},
		},
		{
			name: "success from Internal",
			err: exception.CustomError{
				ErrorCode: exception.ErrInDatastore,
				Value:     test.ErrMock,
			},
			expect: func() error {
				br := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: test.ErrMock.Error(),
				}
				st := status.New(codes.Internal, codes.Internal.String())
				dt, _ := st.WithDetails(br)
				return dt.Err()
			},
		},
		{
			name: "success from Unknown",
			err: exception.CustomError{
				ErrorCode: exception.ErrUnknown,
				Value:     test.ErrMock,
			},
			expect: func() error {
				br := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: test.ErrMock.Error(),
				}
				st := status.New(codes.Unknown, codes.Unknown.String())
				dt, _ := st.WithDetails(br)
				return dt.Err()
			},
		},
		{
			name: "success from Unauthenticated",
			err: exception.CustomError{
				ErrorCode: exception.ErrUnauthorized,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.Unauthenticated, codes.Unauthenticated.String()).Err()
			},
		},
		{
			name: "success from Other gRPC Error",
			err: exception.CustomError{
				ErrorCode: exception.ErrorCode(999),
				Value:     test.ErrMock,
			},
			expect: func() error {
				br := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: test.ErrMock.Error(),
				}
				st := status.New(codes.Unknown, codes.Unknown.String())
				dt, _ := st.WithDetails(br)
				return dt.Err()
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect(), toGRPCError(tt.err))
		})
	}
}

func TestGetError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    error
		expect string
	}{
		{
			name: "success to CustomError type",
			err: exception.CustomError{
				ErrorCode: exception.ErrUnauthorized,
				Value:     test.ErrMock,
			},
			expect: test.ErrMock.Error(),
		},
		{
			name:   "success to other type",
			err:    test.ErrMock,
			expect: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, getError(tt.err))
		})
	}
}

func TestGetErrorCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    error
		expect exception.ErrorCode
	}{
		{
			name: "success to CustomError type",
			err: exception.CustomError{
				ErrorCode: exception.ErrUnauthorized,
				Value:     test.ErrMock,
			},
			expect: exception.ErrUnauthorized,
		},
		{
			name:   "success to other type",
			err:    test.ErrMock,
			expect: exception.ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, getErrorCode(tt.err))
		})
	}
}

func TestGetGRPCErrorCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		code   exception.ErrorCode
		expect codes.Code
	}{
		{
			name:   "InvalidArgument from ErrInvalidRequestValidation",
			code:   exception.ErrInvalidRequestValidation,
			expect: codes.InvalidArgument,
		},
		{
			name:   "DeadlineExceeded from ErrGatewayTimeout",
			code:   exception.ErrGatewayTimeout,
			expect: codes.DeadlineExceeded,
		},
		{
			name:   "NotFound from ErrNotFound",
			code:   exception.ErrNotFound,
			expect: codes.NotFound,
		},
		{
			name:   "AlreadyExists from Conflict",
			code:   exception.ErrConflict,
			expect: codes.AlreadyExists,
		},
		{
			name:   "PermissionDenied from ErrForbidden",
			code:   exception.ErrForbidden,
			expect: codes.PermissionDenied,
		},
		{
			name:   "FailedPrecondition from ErrPreconditionFailed",
			code:   exception.ErrPreconditionFailed,
			expect: codes.FailedPrecondition,
		},
		{
			name:   "Unimplemented from ErrNotImplemented",
			code:   exception.ErrNotImplemented,
			expect: codes.Unimplemented,
		},
		{
			name:   "Internal from ErrInvalidDomainValidation",
			code:   exception.ErrInvalidDomainValidation,
			expect: codes.Internal,
		},
		{
			name:   "Unauthenticated from ErrUnauthorized",
			code:   exception.ErrUnauthorized,
			expect: codes.Unauthenticated,
		},
		{
			name:   "Unknown from non exist error code",
			code:   exception.ErrorCode(999),
			expect: codes.Unknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, getGRPCErrorCode(tt.code))
		})
	}
}

func TestGetValidationErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    error
		expect []*errdetails.BadRequest_FieldViolation
	}{
		{
			name: "success to CustomError type",
			err: exception.CustomError{
				ErrorCode: exception.ErrInvalidRequestValidation,
				Value:     test.ErrMock,
				ValidationErrors: []*exception.ValidationError{
					{
						Field:   "id",
						Message: exception.RequiredMessage,
					},
				},
			},
			expect: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "id",
					Description: exception.RequiredMessage,
				},
			},
		},
		{
			name:   "success to other type",
			err:    test.ErrMock,
			expect: []*errdetails.BadRequest_FieldViolation{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, getValidationErrors(tt.err))
		})
	}
}
