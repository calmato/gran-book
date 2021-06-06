package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestUserDomainValidation_User(t *testing.T) {
	type args struct {
		user *user.User
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				user: &user.User{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "test-user",
					Gender:           0,
					Email:            "test-user@calmato.com",
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

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetUIDByEmail(ctx, tc.args.user.Email).Return(tc.args.user.ID, nil)

		t.Run(name, func(t *testing.T) {
			target := NewUserDomainValidation(urm)

			got := target.User(ctx, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
			}
		})
	}
}

func TestUserDomainValidation_Relationship(t *testing.T) {
	type args struct {
		relationship *user.Relationship
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
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

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().
			GetRelationshipIDByUID(ctx, tc.args.relationship.FollowID, tc.args.relationship.FollowerID).
			Return(tc.args.relationship.ID, nil)

		t.Run(name, func(t *testing.T) {
			target := NewUserDomainValidation(urm)

			got := target.Relationship(ctx, tc.args.relationship)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
			}
		})
	}
}
