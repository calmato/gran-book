package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/golang/mock/gomock"
)

func TestUser_ListUserFollow(t *testing.T) {
	t.Parallel()

	follows := make([]*user.Follow, 2)
	follows[0] = testFollow("11111111-1111-1111-1111-111111111111")
	follows[1] = testFollow("22222222-2222-2222-2222-222222222222")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					ListFollow(gomock.Any(), &user.ListFollowRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(&user.FollowListResponse{
						Follows: follows,
						Total:   2,
						Limit:   100,
						Offset:  0,
					}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewFollowListResponse(entity.NewFollows(follows), 100, 0, 2),
			},
		},
		{
			name:  "failed to invalid limit query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?limit=1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?offset=0.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list follow",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					ListFollow(gomock.Any(), &user.ListFollowRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/follows" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUser_ListUserFollower(t *testing.T) {
	t.Parallel()

	followers := make([]*user.Follower, 2)
	followers[0] = testFollower("11111111-1111-1111-1111-111111111111")
	followers[1] = testFollower("22222222-2222-2222-2222-222222222222")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					ListFollower(gomock.Any(), &user.ListFollowerRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(&user.FollowerListResponse{
						Followers: followers,
						Total:     2,
						Limit:     100,
						Offset:    0,
					}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewFollowerListResponse(entity.NewFollowers(followers), 100, 0, 2),
			},
		},
		{
			name:  "failed to invalid limit query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?limit=1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?offset=0.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list follow",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					ListFollower(gomock.Any(), &user.ListFollowerRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/followers" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUser_GetUserProfile(t *testing.T) {
	t.Parallel()

	profile1 := testProfile("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					GetUserProfile(gomock.Any(), &user.GetUserProfileRequest{UserId: "00000000-0000-0000-0000-000000000000"}).
					Return(&user.UserProfileResponse{Profile: profile1}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewUserProfileResponse(entity.NewUserProfile(profile1)),
			},
		},
		{
			name: "failed to list follow",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					GetUserProfile(gomock.Any(), &user.GetUserProfileRequest{UserId: "00000000-0000-0000-0000-000000000000"}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/profile"
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUser_UserFollow(t *testing.T) {
	t.Parallel()

	profile1 := testProfile("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					Follow(gomock.Any(), &user.FollowRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						FollowerId: "11111111-1111-1111-1111-111111111111",
					}).
					Return(&user.UserProfileResponse{Profile: profile1}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewUserProfileResponse(entity.NewUserProfile(profile1)),
			},
		},
		{
			name: "failed to list follow",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					Follow(gomock.Any(), &user.FollowRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						FollowerId: "11111111-1111-1111-1111-111111111111",
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/follow/11111111-1111-1111-1111-111111111111"
			req := test.NewHTTPRequest(t, http.MethodPost, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUser_UserUnfollow(t *testing.T) {
	t.Parallel()

	profile1 := testProfile("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					Unfollow(gomock.Any(), &user.UnfollowRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						FollowerId: "11111111-1111-1111-1111-111111111111",
					}).
					Return(&user.UserProfileResponse{Profile: profile1}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewUserProfileResponse(entity.NewUserProfile(profile1)),
			},
		},
		{
			name: "failed to list follow",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.UserService.EXPECT().
					Unfollow(gomock.Any(), &user.UnfollowRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						FollowerId: "11111111-1111-1111-1111-111111111111",
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/follow/11111111-1111-1111-1111-111111111111"
			req := test.NewHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testUser(id string) *user.User {
	now := datetime.FormatTime(test.TimeMock)
	return &user.User{
		Id:               id,
		Username:         "テストユーザー",
		Gender:           user.Gender_GENDER_MAN,
		Email:            "test-user@calmato.jp",
		PhoneNumber:      "000-0000-0000",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		LastName:         "テスト",
		FirstName:        "ユーザー",
		LastNameKana:     "てすと",
		FirstNameKana:    "ゆーざー",
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}

func testProfile(id string) *user.UserProfile {
	return &user.UserProfile{
		Id:               id,
		Username:         "テストユーザー",
		ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollow:         false,
		IsFollower:       true,
		FollowCount:      3,
		FollowerCount:    8,
	}
}

func testFollow(id string) *user.Follow {
	return &user.Follow{
		Id:               id,
		Username:         "テストユーザー",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollow:         true,
	}
}

func testFollower(id string) *user.Follower {
	return &user.Follower{
		Id:               id,
		Username:         "テストユーザー",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollow:         false,
	}
}
