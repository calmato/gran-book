package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestBookRequestValidation_Book(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Book
		Expected bool
	}{
		"ok": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.Book{
				Title:        "",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.Book{
				Title:        strings.Repeat("x", 65),
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  strings.Repeat("x", 1601),
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_isbn_required": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_isbn_max": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         strings.Repeat("x", 33),
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_version_required": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_publisher_max": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    strings.Repeat("x", 33),
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_authors_required": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{""},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_authors_max": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{strings.Repeat("x", 33)},
				Categories:   []string{"Comics & Graphic Novels"},
			},
			Expected: false,
		},
		"ng_categories_required": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{""},
			},
			Expected: false,
		},
		"ng_categories_max": {
			Input: &input.Book{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors:      []string{"テスト著者"},
				Categories:   []string{strings.Repeat("x", 33)},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.Book(tc.Input)
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

func TestBookRequestValidation_Bookshelf(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Bookshelf
		Expected bool
	}{
		"ok": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: true,
		},
		"ng_userId_required": {
			Input: &input.Bookshelf{
				UserID:     "",
				BookID:     1,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
		"ng_bookId_required": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     0,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
		"ng_bookId_greater_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     0,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
		"ng_status_greater_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     -1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
		"ng_status_less_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     6,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
		"ng_impression_max": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     1,
				Impression: strings.Repeat("x", 1001),
				ReadOn:     "2020-01-01",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.Bookshelf(tc.Input)
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

func TestBookRequestValidation_CreateAndUpdateBooks(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAndUpdateBooks
		Expected bool
	}{
		"ok": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        strings.Repeat("x", 65),
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  strings.Repeat("x", 1601),
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_isbn_required": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_isbn_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         strings.Repeat("x", 33),
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_version_required": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_publisher_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    strings.Repeat("x", 33),
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_authors_required": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{""},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_authors_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{strings.Repeat("x", 33)},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: false,
		},
		"ng_categories_required": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{""},
					},
				},
			},
			Expected: false,
		},
		"ng_categories_max": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{strings.Repeat("x", 33)},
					},
				},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.CreateAndUpdateBooks(tc.Input)
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
