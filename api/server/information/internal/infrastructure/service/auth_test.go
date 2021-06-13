package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	mock_auth "github.com/calmato/gran-book/api/server/information/mock/domain/auth"
	"github.com/golang/mock/gomock"
)

func TestAuthService_Authentication(t *testing.T) {
	type want struct {
		uid string
		err error
	}

	testCases := map[string]struct {
		want want
	}{
		"ok": {
			want: want{
				uid: "00000000-0000-0000-0000-000000000000",
				err: nil,
			},
		},
		"ng_unauthorized": {
			want: want{
				uid: "",
				err: exception.Unauthorized.New(nil),
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arm := mock_auth.NewMockRepository(ctrl)
		arm.EXPECT().Authentication(ctx).Return(tc.want.uid, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAuthService(arm)

			got, err := target.Authentication(ctx)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(got, tc.want.uid) {
				t.Fatalf("want %#v, but %#v", tc.want.uid, got)
				return
			}
		})
	}
}
