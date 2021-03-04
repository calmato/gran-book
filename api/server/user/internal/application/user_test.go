package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestUserApplication_GetProfile(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		CUID     string
		Expected struct {
			User   *user.User
			Output *output.GetUserProfile
			Error  error
		}
	}{
		"ok": {
			UID:  "11111111-1111-1111-1111-111111111111",
			CUID: "00000000-0000-0000-0000-000000000000",
			Expected: struct {
				User   *user.User
				Output *output.GetUserProfile
				Error  error
			}{
				User: &user.User{
					ID: "11111111-1111-1111-1111-111111111111",
				},
				Output: &output.GetUserProfile{
					IsFollow:       true,
					IsFollower:     false,
					FollowsTotal:   3,
					FollowersTotal: 5,
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().
			ListFriendsCount(ctx, tc.Expected.User).
			Return(tc.Expected.Output.FollowsTotal, tc.Expected.Output.FollowersTotal, tc.Expected.Error)
		usm.EXPECT().
			IsFriend(ctx, tc.Expected.User, tc.CUID).
			Return(tc.Expected.Output.IsFollow, tc.Expected.Output.IsFollower, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.GetProfile(ctx, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(user, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, user)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}
