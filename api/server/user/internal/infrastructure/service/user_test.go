package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestUserService_Authentication(t *testing.T) {
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Authentication(ctx).Return(tc.Expected.UID, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

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

func TestUserService_List(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Query    *domain.ListQuery
		Expected struct {
			Users []*user.User
			Total int64
			Error error
		}
	}{
		"ok": {
			Query: &domain.ListQuery{
				Limit:      100,
				Offset:     0,
				Order:      nil,
				Conditions: []*domain.QueryCondition{},
			},
			Expected: struct {
				Users []*user.User
				Total int64
				Error error
			}{
				Users: []*user.User{
					{
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
						Activated:        true,
						CreatedAt:        current,
						UpdatedAt:        current,
					},
				},
				Total: 1,
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().List(ctx, tc.Query).Return(tc.Expected.Users, tc.Expected.Error)
		urm.EXPECT().ListCount(ctx, tc.Query).Return(tc.Expected.Total, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			users, total, err := target.List(ctx, tc.Query)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(users, tc.Expected.Users) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Users, users)
				return
			}

			if !reflect.DeepEqual(total, tc.Expected.Total) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Total, total)
				return
			}
		})
	}
}

func TestUserService_ListFriendsCount(t *testing.T) {
	testCases := map[string]struct {
		User     *user.User
		Expected struct {
			FollowsCount   int64
			FollowersCount int64
			Error          error
		}
	}{
		"ok": {
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			Expected: struct {
				FollowsCount   int64
				FollowersCount int64
				Error          error
			}{
				FollowsCount:   3,
				FollowersCount: 5,
				Error:          nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		followsQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follow_id",
					Operator: "==",
					Value:    tc.User.ID,
				},
			},
		}

		followersQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follower_id",
					Operator: "==",
					Value:    tc.User.ID,
				},
			},
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollowsCount(ctx, followsQuery).Return(tc.Expected.FollowsCount, tc.Expected.Error)
		urm.EXPECT().ListFollowersCount(ctx, followersQuery).Return(tc.Expected.FollowersCount, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			followsCount, followersCount, err := target.ListFriendsCount(ctx, tc.User)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(followsCount, tc.Expected.FollowsCount) {
				t.Fatalf("want %#v, but %#v", tc.Expected.FollowsCount, followsCount)
				return
			}

			if !reflect.DeepEqual(followersCount, tc.Expected.FollowersCount) {
				t.Fatalf("want %#v, but %#v", tc.Expected.FollowersCount, followersCount)
				return
			}
		})
	}
}

func TestUserService_Show(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UID      string
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			UID: "00000000-0000-0000-0000-000000000000",
			Expected: struct {
				User  *user.User
				Error error
			}{
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
					Activated:        true,
					CreatedAt:        current,
					UpdatedAt:        current,
				},
				Error: nil,
			},
		},
		"ng_notfound": {
			UID: "",
			Expected: struct {
				User  *user.User
				Error error
			}{
				User:  nil,
				Error: exception.NotFound.New(nil),
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.Show(ctx, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, got)
				return
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	testCases := map[string]struct {
		User     *user.User
		Expected error
	}{
		"ok": {
			User: &user.User{
				ID:               "",
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
				Activated:        true,
				CreatedAt:        time.Time{},
				UpdatedAt:        time.Time{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().User(ctx, tc.User).Return(nil)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Create(ctx, tc.User).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Create(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.User.ID == "" {
				t.Fatal("User.ID must be not null")
				return
			}

			if tc.User.CreatedAt.IsZero() {
				t.Fatal("User.CreatedAt must be not null")
				return
			}

			if tc.User.UpdatedAt.IsZero() {
				t.Fatal("User.UpdatedAt must be not null")
				return
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
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
				Activated:        true,
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

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().User(ctx, tc.User).Return(nil)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Update(ctx, tc.User).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Update(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.User.UpdatedAt == current {
				t.Fatal("User.UpdatedAt must be changed")
				return
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		Password string
		Expected error
	}{
		"ok": {
			UID:      "00000000-0000-0000-0000-000000000000",
			Password: "12345678",
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().UpdatePassword(ctx, tc.UID, tc.Password).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.UpdatePassword(ctx, tc.UID, tc.Password)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestUserService_UpdateThumbnail(t *testing.T) {
	testCases := map[string]struct {
		UID       string
		Thumbnail []byte
		Expected  struct {
			ThumbailURL string
			Error       error
		}
	}{
		"ok": {
			UID:       "00000000-0000-0000-0000-000000000000",
			Thumbnail: []byte{},
			Expected: struct {
				ThumbailURL string
				Error       error
			}{
				ThumbailURL: "https://calmato.com/user_thumbnails",
				Error:       nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)
		uum.EXPECT().Thumbnail(ctx, tc.UID, tc.Thumbnail).Return(tc.Expected.ThumbailURL, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.UploadThumbnail(ctx, tc.UID, tc.Thumbnail)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.ThumbailURL) {
				t.Fatalf("want %#v, but %#v", tc.Expected.ThumbailURL, got)
				return
			}
		})
	}
}

func TestUserService_IsFriend(t *testing.T) {
	testCases := map[string]struct {
		User     *user.User
		CUID     string
		Expected struct {
			IsFollow   bool
			IsFollower bool
			Error      error
		}
	}{
		"ok": {
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			CUID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				IsFollow   bool
				IsFollower bool
				Error      error
			}{
				IsFollow:   true,
				IsFollower: false,
				Error:      nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		followsQuery := &domain.ListQuery{
			Limit: 1,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follow_id",
					Operator: "==",
					Value:    tc.User.ID,
				},
				{
					Field:    "follower_id",
					Operator: "==",
					Value:    tc.CUID,
				},
			},
		}

		followersQuery := &domain.ListQuery{
			Limit: 1,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follower_id",
					Operator: "==",
					Value:    tc.User.ID,
				},
				{
					Field:    "follow_id",
					Operator: "==",
					Value:    tc.CUID,
				},
			},
		}

		follows := []*user.User{}
		followers := []*user.User{}

		if tc.Expected.IsFollow {
			followers = append(follows, &user.User{ID: tc.CUID})
		}

		if tc.Expected.IsFollower {
			follows = append(followers, &user.User{ID: tc.CUID})
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollows(ctx, followsQuery).Return(follows, tc.Expected.Error)
		urm.EXPECT().ListFollowers(ctx, followersQuery).Return(followers, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			isFollow, isFollower, err := target.IsFriend(ctx, tc.User, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(isFollow, tc.Expected.IsFollow) {
				t.Fatalf("want %#v, but %#v", tc.Expected.IsFollow, isFollow)
				return
			}

			if !reflect.DeepEqual(isFollower, tc.Expected.IsFollower) {
				t.Fatalf("want %#v, but %#v", tc.Expected.IsFollower, isFollower)
				return
			}
		})
	}
}
