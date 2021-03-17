package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestBookRequestValidation_CreateBook(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateBook
		Expected bool
	}{
		"ok": {
			Input: &input.CreateBook{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.CreateBook{
				Title:        "",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.CreateBook{
				Title:        strings.Repeat("x", 33),
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: false,
		},
		"ng_description": {
			Input: &input.CreateBook{
				Title:        "テスト書籍",
				Description:  strings.Repeat("x", 1001),
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: false,
		},
		"ng_isbn_required": {
			Input: &input.CreateBook{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: false,
		},
		"ng_publisher_max": {
			Input: &input.CreateBook{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    strings.Repeat("x", 33),
				PublishedOn:  "2021-01-01",
				Authors: []*input.CreateBookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.CreateBookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.CreateBook(tc.Input)
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
