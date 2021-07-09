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
	type want struct {
		uid string
		err error
	}

	testCases := map[string]struct {
		want want
	}{
		"ok": {
			want: want{
				uid: "00000000-0000-0000-0000-000000000000",
				err: nil,
			},
		},
		"ng_unauthorized": {
			want: want{
				uid: "",
				err: exception.Unauthorized.New(nil),
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Authentication(ctx).Return(tc.want.uid, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

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

func TestUserService_List(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		users []*user.User
		err   error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				users: []*user.User{
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
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().List(ctx, tc.args.query).Return(tc.want.users, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			us, err := target.List(ctx, tc.args.query)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(us, tc.want.users) {
				t.Fatalf("want %#v, but %#v", tc.want.users, us)
				return
			}
		})
	}
}

func TestUserService_ListFollow(t *testing.T) {
	type args struct {
		query *domain.ListQuery
		uid   string
	}
	type want struct {
		follows []*user.Follow
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
				uid: "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				follows: []*user.Follow{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
						IsFollow:         false,
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollow(ctx, tc.args.query).Return(tc.want.follows, tc.want.err)
		urm.EXPECT().ListFollowerID(ctx, gomock.Any()).Return([]string{}, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			fs, err := target.ListFollow(ctx, tc.args.query, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(fs, tc.want.follows) {
				t.Fatalf("want %#v, but %#v", tc.want.follows, fs)
				return
			}
		})
	}
}

func TestUserService_ListFollower(t *testing.T) {
	type args struct {
		query *domain.ListQuery
		uid   string
	}
	type want struct {
		followers []*user.Follower
		err       error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
				uid: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				followers: []*user.Follower{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
						IsFollow:         false,
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListFollower(ctx, tc.args.query).Return(tc.want.followers, tc.want.err)
		urm.EXPECT().ListFollowerID(ctx, gomock.Any()).Return([]string{}, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			fs, err := target.ListFollower(ctx, tc.args.query, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(fs, tc.want.followers) {
				t.Fatalf("want %#v, but %#v", tc.want.followers, fs)
				return
			}
		})
	}
}

func TestUserService_ListCount(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		count int
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				count: 1,
				err:   nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListCount(ctx, tc.args.query).Return(tc.want.count, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			count, err := target.ListCount(ctx, tc.args.query)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(count, tc.want.count) {
				t.Fatalf("want %#v, but %#v", tc.want.count, count)
				return
			}
		})
	}
}

func TestUserService_ListFriendCount(t *testing.T) {
	type args struct {
		uid string
	}
	type want struct {
		followCount   int
		followerCount int
		err           error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				followCount:   1,
				followerCount: 3,
				err:           nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		followQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follow_id",
					Operator: "==",
					Value:    tc.args.uid,
				},
			},
		}

		followerQuery := &domain.ListQuery{
			Conditions: []*domain.QueryCondition{
				{
					Field:    "follower_id",
					Operator: "==",
					Value:    tc.args.uid,
				},
			},
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ListRelationshipCount(ctx, followQuery).Return(tc.want.followCount, tc.want.err)
		urm.EXPECT().ListRelationshipCount(ctx, followerQuery).Return(tc.want.followerCount, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			followCount, followerCount, err := target.ListFriendCount(ctx, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(followCount, tc.want.followCount) {
				t.Fatalf("want %#v, but %#v", tc.want.followCount, followCount)
				return
			}

			if !reflect.DeepEqual(followerCount, tc.want.followerCount) {
				t.Fatalf("want %#v, but %#v", tc.want.followerCount, followerCount)
				return
			}
		})
	}
}

func TestUserService_Show(t *testing.T) {
	type args struct {
		uid string
	}
	type want struct {
		user *user.User
		err  error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
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
				err: nil,
			},
		},
		"ng_notfound": {
			args: args{
				uid: "",
			},
			want: want{
				user: nil,
				err:  exception.NotFound.New(nil),
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			u, err := target.Show(ctx, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(u, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, u)
				return
			}
		})
	}
}

func TestUserService_ShowRelationship(t *testing.T) {
	type args struct {
		id int
	}
	type want struct {
		relationship *user.Relationship
		err          error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				id: 1,
			},
			want: want{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  current,
					UpdatedAt:  current,
				},
				err: nil,
			},
		},
		"ng_notfound": {
			args: args{
				id: 1,
			},
			want: want{
				relationship: nil,
				err:          exception.NotFound.New(nil),
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().ShowRelationship(ctx, tc.args.id).Return(tc.want.relationship, tc.want.err)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			r, err := target.ShowRelationship(ctx, tc.args.id)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(r, tc.want.relationship) {
				t.Fatalf("want %#v, but %#v", tc.want.relationship, r)
				return
			}
		})
	}
}

func TestUserService_ShowRelationshipByUID(t *testing.T) {
	type args struct {
		followID   string
		followerID string
	}
	type want struct {
		relationship *user.Relationship
		err          error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				followID:   "00000000-0000-0000-0000-000000000000",
				followerID: "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				relationship: &user.Relationship{
					ID:         1,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  current,
					UpdatedAt:  current,
				},
				err: nil,
			},
		},
		"ng_notfound": {
			args: args{
				followID:   "00000000-0000-0000-0000-000000000000",
				followerID: "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				relationship: nil,
				err:          exception.NotFound.New(nil),
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var fid int
		if tc.want.relationship != nil {
			fid = tc.want.relationship.ID
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.args.followID, tc.args.followerID).Return(fid, tc.want.err)
		if fid != 0 {
			urm.EXPECT().ShowRelationship(ctx, fid).Return(tc.want.relationship, tc.want.err)
		}

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			r, err := target.ShowRelationshipByUID(ctx, tc.args.followID, tc.args.followerID)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(r, tc.want.relationship) {
				t.Fatalf("want %#v, but %#v", tc.want.relationship, r)
				return
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		user *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				user: &user.User{
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Create(ctx, tc.args.user).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Create(ctx, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.user.ID == "" {
				t.Fatal("User.ID must be not null")
				return
			}

			if tc.args.user.CreatedAt.IsZero() {
				t.Fatal("User.CreatedAt must be not null")
				return
			}

			if tc.args.user.UpdatedAt.IsZero() {
				t.Fatal("User.UpdatedAt must be not null")
				return
			}
		})
	}
}

func TestUserService_CreateRelationship(t *testing.T) {
	type args struct {
		relationship *user.Relationship
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				relationship: &user.Relationship{
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().CreateRelationship(ctx, tc.args.relationship).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.CreateRelationship(ctx, tc.args.relationship)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.relationship.CreatedAt.IsZero() {
				t.Fatal("Relationship.CreatedAt must be not null")
				return
			}

			if tc.args.relationship.UpdatedAt.IsZero() {
				t.Fatal("Relationship.UpdatedAt must be not null")
				return
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
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

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Update(ctx, tc.args.user).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Update(ctx, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.user.UpdatedAt == current {
				t.Fatal("User.UpdatedAt must be changed")
				return
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	type args struct {
		uid      string
		password string
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				uid:      "00000000-0000-0000-0000-000000000000",
				password: "12345678",
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().UpdatePassword(ctx, tc.args.uid, tc.args.password).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.UpdatePassword(ctx, tc.args.uid, tc.args.password)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	type args struct {
		uid string
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				uid: "00000000-0000-0000-0000-000000000000",
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Delete(ctx, tc.args.uid).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Delete(ctx, tc.args.uid)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestUserService_DeleteRelationship(t *testing.T) {
	type args struct {
		id int
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				id: 1,
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().DeleteRelationship(ctx, tc.args.id).Return(tc.want)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.DeleteRelationship(ctx, tc.args.id)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestUserService_UpdateThumbnail(t *testing.T) {
	type args struct {
		uid       string
		thumbnail []byte
	}
	type want struct {
		thumbnailURL string
		err          error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:       "00000000-0000-0000-0000-000000000000",
				thumbnail: []byte{},
			},
			want: want{
				thumbnailURL: "https://calmato.com/user_thumbnails",
				err:          nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)
		uum.EXPECT().Thumbnail(ctx, tc.args.uid, tc.args.thumbnail).Return(tc.want.thumbnailURL, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			thumbnailURL, err := target.UploadThumbnail(ctx, tc.args.uid, tc.args.thumbnail)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(thumbnailURL, tc.want.thumbnailURL) {
				t.Fatalf("want %#v, but %#v", tc.want.thumbnailURL, thumbnailURL)
				return
			}
		})
	}
}

func TestUserService_IsFriend(t *testing.T) {
	type args struct {
		friendID string
		uid      string
	}
	type want struct {
		isFollow   bool
		isFollower bool
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				friendID: "00000000-0000-0000-0000-000000000000",
				uid:      "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				isFollow:   true,
				isFollower: true,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var followID int
		var followerID int

		if tc.want.isFollow {
			followID = 1
		}

		if tc.want.isFollower {
			followerID = 1
		}

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.args.uid, tc.args.friendID).Return(followID, nil)
		urm.EXPECT().GetRelationshipIDByUID(ctx, tc.args.friendID, tc.args.uid).Return(followerID, nil)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			isFollow, isFollower := target.IsFriend(ctx, tc.args.friendID, tc.args.uid)
			if !reflect.DeepEqual(isFollow, tc.want.isFollow) {
				t.Fatalf("want %#v, but %#v", tc.want.isFollow, isFollow)
				return
			}

			if !reflect.DeepEqual(isFollower, tc.want.isFollower) {
				t.Fatalf("want %#v, but %#v", tc.want.isFollower, isFollower)
				return
			}
		})
	}
}

func TestUserService_Validation(t *testing.T) {
	type args struct {
		user *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				user: &user.User{
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
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().User(ctx, tc.args.user).Return(tc.want)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Validation(ctx, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestUserService_ValidationRelationship(t *testing.T) {
	type args struct {
		relationship *user.Relationship
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				relationship: &user.Relationship{
					ID:         0,
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
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

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().Relationship(ctx, tc.args.relationship).Return(tc.want)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(name, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.ValidationRelationship(ctx, tc.args.relationship)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}
