package handler

import (
	"context"
	"net/http"
	// "strings"
	// "encoding/json"
	"log"
	"io"
	"fmt"

	// "google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type errorResponse struct {
	Code codes.Code `json:"code"`
	Message string `json:"message"`
	Errors []*errorResponseDetail `json:"errors"`
}

type errorResponseDetail struct {
	Field string `json:"field"`
	Reason string `json:"reason"`
	Message string `json:"message"`
}

const (
	fallback = `{"error": "failed to marshal error message"}`
)

// CustomHTTPError ー クライアントへ返すエラーエスポンスの整形用
func CustomHTTPError(
	ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler,
	w http.ResponseWriter, _ *http.Request, err error,
) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	// Set response headers
	contentType := marshaler.ContentType(nil)
	w.Header().Set("Content-Type", contentType)

	// Create response body
	res := &errorResponse{
		Code: s.Code(),
		Message: s.Message(),
		Errors: []*errorResponseDetail{},
	}

	for _, detail := range s.Details() {
		switch v := detail.(type) {
		case *errdetails.BadRequest:
			if v.FieldViolations == nil {
				continue
			}

			res.Errors = append(res.Errors, getErrorResponseDetails(v.FieldViolations)...)
		}
	}

	// Convert to custom error response
	buf, err := marshaler.Marshal(res)
	if err != nil {
		log.Printf("Failed to marshal error message %q: %v", res, err)

		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			log.Printf("Failed to write response: %v", err)
		}

		return
	}

	st := runtime.HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err = w.Write(buf); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func getErrorResponseDetails(fvs []*errdetails.BadRequest_FieldViolation) []*errorResponseDetail {
	ds := make([]*errorResponseDetail, len(fvs))
	for i, fv := range fvs {
		ds[i] = &errorResponseDetail{
				Field: fv.Field,
				Reason: fv.Description,
				Message: fmt.Sprintf("%s %s", fv.Field, fv.Description),
		}
	}

	return ds
}
