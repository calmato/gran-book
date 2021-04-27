package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestUserRequestValidation_ListUser(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListUser
		Expected bool
	}{
		"ok": {
			Input: &input.ListUser{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.ListUser{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.ListUser{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.ListUser{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.ListUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.ListUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewUserRequestValidation()

			got := target.ListUser(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestUserRequestValidation_ListUserByUserIDs(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListUserByUserIDs
		Expected bool
	}{
		"ok": {
			Input: &input.ListUserByUserIDs{
				UserIDs: []string{
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
					"78b10377-4e3e-48fd-892c-36bcd7c91f3b",
				},
			},
			Expected: true,
		},
		"ng_userIds_unique": {
			Input: &input.ListUserByUserIDs{
				UserIDs: []string{
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
				},
			},
			Expected: false,
		},
		"ng_userIds_required": {
			Input: &input.ListUserByUserIDs{
				UserIDs: []string{
					"",
					"78b10377-4e3e-48fd-892c-36bcd7c91f3b",
				},
			},
			Expected: false,
		},
		"ng_userIds_max": {
			Input: &input.ListUserByUserIDs{
				UserIDs: []string{
					strings.Repeat("x", 37),
					"78b10377-4e3e-48fd-892c-36bcd7c91f3b",
				},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewUserRequestValidation()

			got := target.ListUserByUserIDs(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestUserRequestValidation_ListFollow(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListFollow
		Expected bool
	}{
		"ok": {
			Input: &input.ListFollow{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.ListFollow{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.ListFollow{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.ListFollow{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.ListFollow{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.ListFollow{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewUserRequestValidation()

			got := target.ListFollow(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestUserRequestValidation_ListFollower(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListFollower
		Expected bool
	}{
		"ok": {
			Input: &input.ListFollower{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.ListFollower{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.ListFollower{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.ListFollower{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.ListFollower{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.ListFollower{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewUserRequestValidation()

			got := target.ListFollower(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestUserRequestValidation_SearchUser(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.SearchUser
		Expected bool
	}{
		"ok": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
				Field:     "email",
				Value:     "test",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.SearchUser{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.SearchUser{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_field_required": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_value_required": {
			Input: &input.SearchUser{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewUserRequestValidation()

			got := target.SearchUser(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}
