package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/domain/inquiry"
	mock_inquiry "github.com/calmato/gran-book/api/server/information/mock/domain/inquiry"
	"github.com/golang/mock/gomock"
)

func TestInquiryService_Create(t *testing.T) {
	type args struct {
		in *input.CreateInquiry
	}

	type want struct {
		inquiry *inquiry.Inquiry
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{},
			want: want{},
		},
		// TODO: 失敗時のテストケース記載
		"ng": {
			args: args{},
			want: want{},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ism := mock_inquiry.NewMockRepository(ctrl)
		ism.EXPECT().Create(ctx, tc.args.in).Return(tc.want.inquiry, tc.want.err)

		// Add dependencies for structure fields.

		t.Run(name, func(t *testing.T) {
			target := NewInquiryService(ism)

			err := target.Create(ctx, tc.want.inquiry)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

		})
	}
}
