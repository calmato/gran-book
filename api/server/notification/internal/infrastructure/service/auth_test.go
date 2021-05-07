package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/notification/internal/domain/exception"
	mock_auth "github.com/calmato/gran-book/api/server/notification/mock/domain/auth"
	"github.com/golang/mock/gomock"
)

func TestAuthService_Authentication(t *testing.T) {
	testCases := map[string]struct {
		Expected struct {
			UID   string
			Error error
		}
	}{
		"ok": {
			Expected: struct {
				UID   string
				Error error
			}{
				UID:   "00000000-0000-0000-0000-000000000000",
				Error: nil,
			},
		},
		"ng_unauthorized": {
			Expected: struct {
				UID   string
				Error error
			}{
				UID:   "",
				Error: exception.Unauthorized.New(nil),
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arm := mock_auth.NewMockRepository(ctrl)
		arm.EXPECT().Authentication(ctx).Return(tc.Expected.UID, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAuthService(arm)

			got, err := target.Authentication(ctx)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.UID) {
				t.Fatalf("want %#v, but %#v", tc.Expected.UID, got)
				return
			}
		})
	}
}
