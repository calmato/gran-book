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
						CreatedAt:        current,
						UpdatedAt:        current,
					},
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().List(ctx, tc.Query).Return(tc.Expected.Users, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.List(ctx, tc.Query)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Users) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Users, got)
				return
			}
		})
	}
}

func TestUserService_ListFollow(t *testing.T) {
	testCases := map[string]struct {
		Query    *domain.ListQuery
		UID      string
		Expected struct {
			Follows []*user.Follow
			Error   error
		}
	}{
		"ok": {
			Query: &domain.ListQuery{
				Limit:      100,
				Offset:     0,
				Order:      nil,
				Conditions: []*domain.QueryCondition{},
			},
			UID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				Follows []*user.Follow
				Error   error
			}{
				Follows: []*user.Follow{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
						IsFollow:         false,
					},
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollow(ctx, tc.Query).Return(tc.Expected.Follows, tc.Expected.Error)
		urm.EXPECT().ListFollowerID(ctx, gomock.Any()).Return([]string{}, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.ListFollow(ctx, tc.Query, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Follows) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Follows, got)
				return
			}
		})
	}
}

func TestUserService_ListFollower(t *testing.T) {
	testCases := map[string]struct {
		Query    *domain.ListQuery
		UID      string
		Expected struct {
			Follower []*user.Follower
			Error    error
		}
	}{
		"ok": {
			Query: &domain.ListQuery{
				Limit:      100,
				Offset:     0,
				Order:      nil,
				Conditions: []*domain.QueryCondition{},
			},
			UID: "00000000-0000-0000-0000-000000000000",
			Expected: struct {
				Follower []*user.Follower
				Error    error
			}{
				Follower: []*user.Follower{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
						IsFollow:         false,
					},
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollower(ctx, tc.Query).Return(tc.Expected.Follower, tc.Expected.Error)
		urm.EXPECT().ListFollowerID(ctx, gomock.Any()).Return([]string{}, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.ListFollower(ctx, tc.Query, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Follower) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Follower, got)
				return
			}
		})
	}
}

func TestUserService_ListCount(t *testing.T) {
	testCases := map[string]struct {
		Query    *domain.ListQuery
		Expected struct {
			Count int
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
				Count int
				Error error
			}{
				Count: 1,
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
		urm.EXPECT().ListCount(ctx, tc.Query).Return(tc.Expected.Count, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.ListCount(ctx, tc.Query)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Count) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Count, got)
				return
			}
		})
	}
}

func TestUserService_ListFriendCount(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		Expected struct {
			FollowCount   int
			FollowerCount int
			Error         error
		}
	}{
		"ok": {
			UID: "00000000-0000-0000-0000-000000000000",
			Expected: struct {
				FollowCount   int
				FollowerCount int
				Error         error
			}{
				FollowCount:   1,
				FollowerCount: 3,
				Error:         nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		followQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follow_id",
					Operator: "==",
					Value:    tc.UID,
				},
			},
		}

		followerQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follower_id",
					Operator: "==",
					Value:    tc.UID,
				},
			},
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListRelationshipCount(ctx, followQuery).Return(tc.Expected.FollowCount, tc.Expected.Error)
		urm.EXPECT().ListRelationshipCount(ctx, followerQuery).Return(tc.Expected.FollowerCount, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			followCount, followerCount, err := target.ListFriendCount(ctx, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(followCount, tc.Expected.FollowCount) {
				t.Fatalf("want %#v, but %#v", tc.Expected.FollowCount, followCount)
				return
			}

			if !reflect.DeepEqual(followerCount, tc.Expected.FollowerCount) {
				t.Fatalf("want %#v, but %#v", tc.Expected.FollowerCount, followerCount)
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

func TestUserService_ShowRelationship(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		ID       int
		Expected struct {
			Relationship *user.Relationship
			Error        error
		}
	}{
		"ok": {
			ID: 1,
			Expected: struct {
				Relationship *user.Relationship
				Error        error
			}{
				Relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  current,
					UpdatedAt:  current,
				},
				Error: nil,
			},
		},
		"ng_notfound": {
			ID: 1,
			Expected: struct {
				Relationship *user.Relationship
				Error        error
			}{
				Relationship: nil,
				Error:        exception.NotFound.New(nil),
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
		urm.EXPECT().ShowRelationship(ctx, tc.ID).Return(tc.Expected.Relationship, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.ShowRelationship(ctx, tc.ID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Relationship) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Relationship, got)
				return
			}
		})
	}
}

func TestUserService_ShowRelationshipByUID(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		FollowID   string
		FollowerID string
		Expected   struct {
			Relationship *user.Relationship
			Error        error
		}
	}{
		"ok": {
			FollowID:   "00000000-0000-0000-0000-000000000000",
			FollowerID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				Relationship *user.Relationship
				Error        error
			}{
				Relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  current,
					UpdatedAt:  current,
				},
				Error: nil,
			},
		},
		"ng_notfound": {
			FollowID:   "00000000-0000-0000-0000-000000000000",
			FollowerID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				Relationship *user.Relationship
				Error        error
			}{
				Relationship: nil,
				Error:        exception.NotFound.New(nil),
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var fid int
		if tc.Expected.Relationship != nil {
			fid = tc.Expected.Relationship.ID
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.FollowID, tc.FollowerID).Return(fid, tc.Expected.Error)
		if fid != 0 {
			urm.EXPECT().ShowRelationship(ctx, fid).Return(tc.Expected.Relationship, tc.Expected.Error)
		}

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.ShowRelationshipByUID(ctx, tc.FollowID, tc.FollowerID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Relationship) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Relationship, got)
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

func TestUserService_CreateRelationship(t *testing.T) {
	testCases := map[string]struct {
		Relationship *user.Relationship
		Expected     error
	}{
		"ok": {
			Relationship: &user.Relationship{
				ID:         0,
				FollowID:   "00000000-0000-0000-0000-000000000000",
				FollowerID: "11111111-1111-1111-1111-111111111111",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
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

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().CreateRelationship(ctx, tc.Relationship).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.CreateRelationship(ctx, tc.Relationship)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.Relationship.CreatedAt.IsZero() {
				t.Fatal("Relationship.CreatedAt must be not null")
				return
			}

			if tc.Relationship.UpdatedAt.IsZero() {
				t.Fatal("Relationship.UpdatedAt must be not null")
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

func TestUserService_Delete(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		Expected error
	}{
		"ok": {
			UID:      "00000000-0000-0000-0000-000000000000",
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
		urm.EXPECT().Delete(ctx, tc.UID).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Delete(ctx, tc.UID)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestUserService_DeleteRelationship(t *testing.T) {
	testCases := map[string]struct {
		ID       int
		Expected error
	}{
		"ok": {
			ID:       1,
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
		urm.EXPECT().DeleteRelationship(ctx, tc.ID).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.DeleteRelationship(ctx, tc.ID)
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
		FriendID string
		UID      string
		Expected struct {
			IsFollow   bool
			IsFollower bool
		}
	}{
		"ok": {
			FriendID: "00000000-0000-0000-0000-000000000000",
			UID:      "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				IsFollow   bool
				IsFollower bool
			}{
				IsFollow:   true,
				IsFollower: true,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var followID int
		var followerID int

		if tc.Expected.IsFollow {
			followID = 1
		}

		if tc.Expected.IsFollower {
			followerID = 1
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.UID, tc.FriendID).Return(followID, nil)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.FriendID, tc.UID).Return(followerID, nil)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			isFollow, isFollower := target.IsFriend(ctx, tc.FriendID, tc.UID)
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

func TestUserService_Validation(t *testing.T) {
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
		uvm.EXPECT().User(ctx, tc.User).Return(tc.Expected)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Validation(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestUserService_ValidationRelationship(t *testing.T) {
	testCases := map[string]struct {
		Relationship *user.Relationship
		Expected     error
	}{
		"ok": {
			Relationship: &user.Relationship{
				ID:         0,
				FollowID:   "00000000-0000-0000-0000-000000000000",
				FollowerID: "11111111-1111-1111-1111-111111111111",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
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
		uvm.EXPECT().Relationship(ctx, tc.Relationship).Return(tc.Expected)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.ValidationRelationship(ctx, tc.Relationship)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}
