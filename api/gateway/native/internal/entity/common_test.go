package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		gender      user.Gender
		key         string
		expect      Gender
		expectName  string
		expectProto user.Gender
	}{
		{
			name:        "success: unnown",
			gender:      user.Gender_GENDER_UNKNOWN,
			key:         "unknown",
			expect:      GenderUnknown,
			expectName:  "unknown",
			expectProto: user.Gender_GENDER_UNKNOWN,
		},
		{
			name:        "success: man",
			gender:      user.Gender_GENDER_MAN,
			key:         "man",
			expect:      GenderMan,
			expectName:  "man",
			expectProto: user.Gender_GENDER_MAN,
		},
		{
			name:        "success: woman",
			gender:      user.Gender_GENDER_WOMAN,
			key:         "woman",
			expect:      GenderWoman,
			expectName:  "woman",
			expectProto: user.Gender_GENDER_WOMAN,
		},
		{
			name:        "success: invalid gender",
			gender:      -1,
			key:         "invalid key",
			expect:      GenderUnknown,
			expectName:  "unknown",
			expectProto: user.Gender_GENDER_UNKNOWN,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewGender(tt.gender)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectName, actual.Name())
			assert.Equal(t, tt.expect, actual.Value(tt.key))
			assert.Equal(t, tt.expectProto, actual.Proto())
		})
	}
}

func TestBookshelfStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		status     book.BookshelfStatus
		key        string
		expect     BookshelfStatus
		expectName string
	}{
		{
			name:       "success: none",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_NONE,
			key:        "none",
			expect:     BookshelfStatusNone,
			expectName: "none",
		},
		{
			name:       "success: read",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_READ,
			key:        "read",
			expect:     BookshelfStatusRead,
			expectName: "read",
		},
		{
			name:       "success: reading",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_READING,
			key:        "reading",
			expect:     BookshelfStatusReading,
			expectName: "reading",
		},
		{
			name:       "success: stacked",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
			key:        "stack",
			expect:     BookshelfStatusStacked,
			expectName: "stack",
		},
		{
			name:       "success: want",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_WANT,
			key:        "want",
			expect:     BookshelfStatusWant,
			expectName: "want",
		},
		{
			name:       "success: release",
			status:     book.BookshelfStatus_BOOKSHELF_STATUS_RELEASE,
			key:        "release",
			expect:     BookshelfStatusRelease,
			expectName: "release",
		},
		{
			name:       "success: invalid status",
			status:     -1,
			key:        "invalid key",
			expect:     BookshelfStatusNone,
			expectName: "none",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelfStatus(tt.status)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectName, actual.Name())
			assert.Equal(t, tt.expect, actual.Value(tt.key))
		})
	}
}
