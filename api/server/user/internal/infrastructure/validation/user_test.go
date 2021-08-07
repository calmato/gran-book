package validation

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserDomainValidation_User(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	current := time.Now().Local()

	type args struct {
		user *user.User
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetUserIDByEmail(ctx, "test-user@calmato.jp").
					Return("00000000-0000-0000-0000-000000000000", nil)
			},
			args: args{
				user: &user.User{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "test-user",
					Gender:           0,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             0,
					ThumbnailURL:     "",
					SelfIntroduction: "",
					LastName:         "テスト",
					FirstName:        "ユーザ",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざ",
					PostalCode:       "000-0000",
					Prefecture:       "東京都",
					City:             "小金井市",
					AddressLine1:     "貫井北町4-1-1",
					AddressLine2:     "",
					InstanceID:       "",
					CreatedAt:        current,
					UpdatedAt:        current,
				},
			},
			want: nil,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserDomainValidation(mocks.UserRepository)

			err := target.User(ctx, tt.args.user)
			if tt.want != nil {
				require.Equal(t, tt.want.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestUserDomainValidation_Relationship(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	current := time.Now().Local()

	type args struct {
		relationship *user.Relationship
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  error
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111").
					Return(1, nil)
			},
			args: args{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  current,
					UpdatedAt:  current,
				},
			},
			want: nil,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserDomainValidation(mocks.UserRepository)

			err := target.Relationship(ctx, tt.args.relationship)
			if tt.want != nil {
				require.Equal(t, tt.want.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
