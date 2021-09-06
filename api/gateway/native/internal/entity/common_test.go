package entity

import (
	"testing"

	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		gender     pb.Gender
		expect     Gender
		expectName string
	}{
		{
			name:       "success: unnown",
			gender:     pb.Gender_GENDER_UNKNOWN,
			expect:     GenderUnknown,
			expectName: "unknown",
		},
		{
			name:       "success: man",
			gender:     pb.Gender_GENDER_MAN,
			expect:     GenderMan,
			expectName: "man",
		},
		{
			name:       "success: woman",
			gender:     pb.Gender_GENDER_WOMAN,
			expect:     GenderWoman,
			expectName: "woman",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewGender(tt.gender)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectName, actual.Name())
			assert.Equal(t, tt.gender, actual.Proto())
		})
	}
}

func TestBookshelfStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		status     pb.BookshelfStatus
		expect     BookshelfStatus
		expectName string
	}{
		{
			name:       "success: none",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_NONE,
			expect:     BookshelfStatusNone,
			expectName: "none",
		},
		{
			name:       "success: read",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_READ,
			expect:     BookshelfStatusRead,
			expectName: "read",
		},
		{
			name:       "success: reading",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
			expect:     BookshelfStatusReading,
			expectName: "reading",
		},
		{
			name:       "success: stacked",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
			expect:     BookshelfStatusStacked,
			expectName: "stack",
		},
		{
			name:       "success: want",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_WANT,
			expect:     BookshelfStatusWant,
			expectName: "want",
		},
		{
			name:       "success: release",
			status:     pb.BookshelfStatus_BOOKSHELF_STATUS_RELEASE,
			expect:     BookshelfStatusRelease,
			expectName: "release",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelfStatus(tt.status)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectName, actual.Name())
		})
	}
}
