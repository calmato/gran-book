package validation

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/internal/user/domain/user"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserDomainValidation_User(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	tests := []struct {
		name   string
		setup  func(context.Context, *testing.T, *test.Mocks)
		args   args
		expect bool
	}{
		{
			name: "success if id does not exists when create",
			setup: func(c context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetUserIDByEmail(ctx, "test-user@calmato.jp").
					Return("", test.ErrMock)
			},
			args: args{
				user: &user.User{
					ID:               "",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: true,
		},
		{
			name: "success if id exists when update",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: true,
		},
		{
			name: "failed if id exists when create",
			setup: func(c context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetUserIDByEmail(ctx, "test-user@calmato.jp").
					Return("00000000-0000-0000-0000-000000000000", nil)
			},
			args: args{
				user: &user.User{
					ID:               "",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: false,
		},
		{
			name: "failed if id is not match when update",
			setup: func(c context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetUserIDByEmail(ctx, "test-user@calmato.jp").
					Return("11111111-1111-1111-1111-111111111111", nil)
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: false,
		},
		{
			name:  "success if email is zero value",
			setup: func(c context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				user: &user.User{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "test-user",
					Gender:           0,
					Email:            "",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserDomainValidation(mocks.UserRepository)

			err := target.User(ctx, tt.args.user)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}

func TestUserDomainValidation_Relationship(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		relationship *user.Relationship
	}
	tests := []struct {
		name   string
		setup  func(context.Context, *testing.T, *test.Mocks)
		args   args
		expect bool
	}{
		{
			name: "success if id does not exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111").
					Return(0, test.ErrMock)
			},
			args: args{
				relationship: &user.Relationship{
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: true,
		},
		{
			name: "success if id does not exists when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111").
					Return(0, test.ErrMock)
			},
			args: args{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: true,
		},
		{
			name: "failed if id exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111").
					Return(1, nil)
			},
			args: args{
				relationship: &user.Relationship{
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: false,
		},
		{
			name: "failed if id is not match when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111").
					Return(2, nil)
			},
			args: args{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: false,
		},
		{
			name:  "success if follow id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: true,
		},
		{
			name: "success if follower id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
			},
			args: args{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserDomainValidation(mocks.UserRepository)

			err := target.Relationship(ctx, tt.args.relationship)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}
