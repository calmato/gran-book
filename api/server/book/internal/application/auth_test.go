package application

import (
	"context"
	"reflect"
	"testing"

	mock_auth "github.com/calmato/gran-book/api/server/book/mock/domain/auth"
	"github.com/golang/mock/gomock"
)

func TestAuthApplication_Authentication(t *testing.T) {
	type want struct {
		uid string
		err error
	}

	testCases := map[string]struct {
		want want
	}{
		"ok": {
			want: want{
				uid:   "00000000-0000-0000-0000-000000000000",
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		asm := mock_auth.NewMockService(ctrl)
		asm.EXPECT().Authentication(ctx).Return(tc.want.uid, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(asm)

			uid, err := target.Authentication(ctx)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(uid, tc.want.uid) {
				t.Fatalf("want %#v, but %#v", tc.want.uid, uid)
				return
			}
		})
	}
}
