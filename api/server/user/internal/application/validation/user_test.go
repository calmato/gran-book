package validation

import (
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

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
