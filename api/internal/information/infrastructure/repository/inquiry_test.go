package repository

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInquiryRepository_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	defer ctrl.Finish()

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		inquiry *inquiry.Inquiry
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				inquiry: &inquiry.Inquiry{
					SenderID:    "00000000-0000-0000-0000-000000000000",
					AdminID:     "11111111-1111-1111-1111-111111111111",
					Subject:     "お問い合わせタイトル",
					Description: "お問い合わせ詳細",
					Email:       "test@calmato.jp",
					IsReplied:   false,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				inquiry: &inquiry.Inquiry{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.InformationDB, inquiryTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewInquiryRepository(mocks.InformationDB, test.Now)
			err = target.Create(ctx, tt.args.inquiry)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}
