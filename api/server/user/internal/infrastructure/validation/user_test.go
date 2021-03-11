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
	current := time.Now()

	testCases := map[string]struct {
		User     *user.User
		Expected error
	}{
		"ok": {
			User: &user.User{
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
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetUIDByEmail(ctx, tc.User.Email).Return(tc.User.ID, nil)

		t.Run(result, func(t *testing.T) {
			target := NewUserDomainValidation(urm)

			got := target.User(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}

func TestUserDomainValidation_Relationship(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Relationship *user.Relationship
		Expected     error
	}{
		"ok": {
			Relationship: &user.Relationship{
				ID:         1,
				FollowID:   "00000000-0000-0000-0000-000000000000",
				FollowerID: "11111111-1111-1111-1111-111111111111",
				CreatedAt:  current,
				UpdatedAt:  current,
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().
			GetRelationshipIDByUID(ctx, tc.Relationship.FollowID, tc.Relationship.FollowerID).
			Return(tc.Relationship.ID, nil)

		t.Run(result, func(t *testing.T) {
			target := NewUserDomainValidation(urm)

			got := target.Relationship(ctx, tc.Relationship)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
			}
		})
	}
}
