package server

import (
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/test"
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
				ErrorCode: exception.InvalidRequestValidation,
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
			name: "success from AlreadyExists",
			err: exception.CustomError{
				ErrorCode: exception.Conflict,
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
			name: "success from Internal",
			err: exception.CustomError{
				ErrorCode: exception.ErrorInDatastore,
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
		{
			name: "success from Other gRPC Error",
			err: exception.CustomError{
				ErrorCode: exception.NotFound,
				Value:     test.ErrMock,
			},
			expect: func() error {
				return status.New(codes.NotFound, codes.NotFound.String()).Err()
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
				ErrorCode: exception.Unauthorized,
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
				ErrorCode: exception.Unauthorized,
				Value:     test.ErrMock,
			},
			expect: exception.Unauthorized,
		},
		{
			name:   "success to other type",
			err:    test.ErrMock,
			expect: exception.Unknown,
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
			name:   "InvalidArgument from InvalidRequestValidation",
			code:   exception.InvalidRequestValidation,
			expect: codes.InvalidArgument,
		},
		{
			name:   "InvalidArgument from UnableConvertBase64",
			code:   exception.UnableConvertBase64,
			expect: codes.InvalidArgument,
		},
		{
			name:   "NotFound from NotFound",
			code:   exception.NotFound,
			expect: codes.NotFound,
		},
		{
			name:   "AlreadyExists from Conflict",
			code:   exception.Conflict,
			expect: codes.AlreadyExists,
		},
		{
			name:   "PermissionDenied from Forbidden",
			code:   exception.Forbidden,
			expect: codes.PermissionDenied,
		},
		{
			name:   "Internal from ErrorInDatastore",
			code:   exception.ErrorInDatastore,
			expect: codes.Internal,
		},
		{
			name:   "Internal from ErrorInStorage",
			code:   exception.ErrorInStorage,
			expect: codes.Internal,
		},
		{
			name:   "Internal from ErrorInOtherAPI",
			code:   exception.ErrorInOtherAPI,
			expect: codes.Internal,
		},
		{
			name:   "Internal from InvalidDomainValidation",
			code:   exception.InvalidDomainValidation,
			expect: codes.Internal,
		},
		{
			name:   "Unauthenticated from Unauthorized",
			code:   exception.Unauthorized,
			expect: codes.Unauthenticated,
		},
		{
			name:   "Unauthenticated from Expired",
			code:   exception.Expired,
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
				ErrorCode: exception.InvalidRequestValidation,
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
