package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		gender user.Gender
		expect Gender
	}{
		{
			name:   "success: unnown",
			gender: user.Gender_GENDER_UNKNOWN,
			expect: GenderUnknown,
		},
		{
			name:   "success: man",
			gender: user.Gender_GENDER_MAN,
			expect: GenderMan,
		},
		{
			name:   "success: woman",
			gender: user.Gender_GENDER_WOMAN,
			expect: GenderWoman,
		},
		{
			name:   "success: invalid gender",
			gender: -1,
			expect: GenderUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewGender(tt.gender)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestGender_ByValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		key    string
		expect Gender
	}{
		{
			name:   "success: unnown",
			key:    "unknown",
			expect: GenderUnknown,
		},
		{
			name:   "success: man",
			key:    "man",
			expect: GenderMan,
		},
		{
			name:   "success: woman",
			key:    "woman",
			expect: GenderWoman,
		},
		{
			name:   "success: invalid gender",
			key:    "invalid key",
			expect: GenderUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewGenderByValue(tt.key)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestGender_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		gender Gender
		expect string
	}{
		{
			name:   "success: unnown",
			gender: GenderUnknown,
			expect: "unknown",
		},
		{
			name:   "success: man",
			gender: GenderMan,
			expect: "man",
		},
		{
			name:   "success: woman",
			gender: GenderWoman,
			expect: "woman",
		},
		{
			name:   "success: invalid gender",
			gender: -1,
			expect: "unknown",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.gender.Name())
		})
	}
}

func TestGender_Proto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		gender Gender
		expect user.Gender
	}{
		{
			name:   "success: unnown",
			gender: GenderUnknown,
			expect: user.Gender_GENDER_UNKNOWN,
		},
		{
			name:   "success: man",
			gender: GenderMan,
			expect: user.Gender_GENDER_MAN,
		},
		{
			name:   "success: woman",
			gender: GenderWoman,
			expect: user.Gender_GENDER_WOMAN,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.gender.Proto())
		})
	}
}

func TestRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   user.Role
		expect Role
	}{
		{
			name:   "success: user",
			role:   user.Role_ROLE_USER,
			expect: RoleUser,
		},
		{
			name:   "success: admin",
			role:   user.Role_ROLE_ADMIN,
			expect: RoleAdmin,
		},
		{
			name:   "success: developer",
			role:   user.Role_ROLE_DEVELOPER,
			expect: RoleDeveloper,
		},
		{
			name:   "success: operator",
			role:   user.Role_ROLE_OPERATOR,
			expect: RoleOperator,
		},
		{
			name:   "success: other role",
			role:   -1,
			expect: RoleUser,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewRole(tt.role)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestRole_ByValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		key    string
		expect Role
	}{
		{
			name:   "success: user",
			key:    "user",
			expect: RoleUser,
		},
		{
			name:   "success: admin",
			key:    "admin",
			expect: RoleAdmin,
		},
		{
			name:   "success: developer",
			key:    "developer",
			expect: RoleDeveloper,
		},
		{
			name:   "success: operator",
			key:    "operator",
			expect: RoleOperator,
		},
		{
			name:   "success: other role",
			key:    "",
			expect: RoleUser,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewRoleByValue(tt.key)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestRole_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   Role
		expect string
	}{
		{
			name:   "success: user",
			role:   RoleUser,
			expect: "user",
		},
		{
			name:   "success: admin",
			role:   RoleAdmin,
			expect: "admin",
		},
		{
			name:   "success: developer",
			role:   RoleDeveloper,
			expect: "developer",
		},
		{
			name:   "success: operator",
			role:   RoleOperator,
			expect: "operator",
		},
		{
			name:   "success: other role",
			role:   -1,
			expect: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.role.Name())
		})
	}
}

func TestRole_Proto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   Role
		expect user.Role
	}{
		{
			name:   "success: user",
			role:   RoleUser,
			expect: user.Role_ROLE_USER,
		},
		{
			name:   "success: admin",
			role:   RoleAdmin,
			expect: user.Role_ROLE_ADMIN,
		},
		{
			name:   "success: developer",
			role:   RoleDeveloper,
			expect: user.Role_ROLE_DEVELOPER,
		},
		{
			name:   "success: operator",
			role:   RoleOperator,
			expect: user.Role_ROLE_OPERATOR,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.role.Proto())
		})
	}
}

func TestBookshelfStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status book.BookshelfStatus
		expect BookshelfStatus
	}{
		{
			name:   "success: none",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_NONE,
			expect: BookshelfStatusNone,
		},
		{
			name:   "success: read",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_READ,
			expect: BookshelfStatusRead,
		},
		{
			name:   "success: reading",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_READING,
			expect: BookshelfStatusReading,
		},
		{
			name:   "success: stacked",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
			expect: BookshelfStatusStacked,
		},
		{
			name:   "success: want",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_WANT,
			expect: BookshelfStatusWant,
		},
		{
			name:   "success: release",
			status: book.BookshelfStatus_BOOKSHELF_STATUS_RELEASE,
			expect: BookshelfStatusRelease,
		},
		{
			name:   "success: invalid status",
			status: -1,
			expect: BookshelfStatusNone,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelfStatus(tt.status)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBookshelfStatus_ByValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		key    string
		expect BookshelfStatus
	}{
		{
			name:   "success: none",
			key:    "none",
			expect: BookshelfStatusNone,
		},
		{
			name:   "success: read",
			key:    "read",
			expect: BookshelfStatusRead,
		},
		{
			name:   "success: reading",
			key:    "reading",
			expect: BookshelfStatusReading,
		},
		{
			name:   "success: stacked",
			key:    "stack",
			expect: BookshelfStatusStacked,
		},
		{
			name:   "success: want",
			key:    "want",
			expect: BookshelfStatusWant,
		},
		{
			name:   "success: release",
			key:    "release",
			expect: BookshelfStatusRelease,
		},
		{
			name:   "success: invalid status",
			key:    "invalid key",
			expect: BookshelfStatusNone,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelfStatusByValue(tt.key)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBookshelfStatus_Name(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status BookshelfStatus
		expect string
	}{
		{
			name:   "success: none",
			status: BookshelfStatusNone,
			expect: "none",
		},
		{
			name:   "success: read",
			status: BookshelfStatusRead,
			expect: "read",
		},
		{
			name:   "success: reading",
			status: BookshelfStatusReading,
			expect: "reading",
		},
		{
			name:   "success: stacked",
			status: BookshelfStatusStacked,
			expect: "stack",
		},
		{
			name:   "success: want",
			status: BookshelfStatusWant,
			expect: "want",
		},
		{
			name:   "success: release",
			status: BookshelfStatusRelease,
			expect: "release",
		},
		{
			name:   "success: invalid status",
			status: -1,
			expect: "none",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.status.Name())
		})
	}
}

func TestBookshelfStatus_Proto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status BookshelfStatus
		expect book.BookshelfStatus
	}{
		{
			name:   "success: none",
			status: BookshelfStatusNone,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_NONE,
		},
		{
			name:   "success: read",
			status: BookshelfStatusRead,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_READ,
		},
		{
			name:   "success: reading",
			status: BookshelfStatusReading,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_READING,
		},
		{
			name:   "success: stacked",
			status: BookshelfStatusStacked,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
		},
		{
			name:   "success: want",
			status: BookshelfStatusWant,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_WANT,
		},
		{
			name:   "success: release",
			status: BookshelfStatusRelease,
			expect: book.BookshelfStatus_BOOKSHELF_STATUS_RELEASE,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.status.Proto())
		})
	}
}
