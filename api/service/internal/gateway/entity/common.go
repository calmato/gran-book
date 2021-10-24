package entity

import (
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
)

// OrderBy - ソート順
type OrderBy int32

const (
	OrderByAsc  OrderBy = 0 // 昇順
	OrderByDesc OrderBy = 1 // 降順
)

var (
	orderByName = map[OrderBy]string{
		0: "asc",
		1: "desc",
	}
	orderByValue = map[string]int32{
		"asc":  0,
		"desc": 1,
	}
)

func NewOrderByByValue(key string) OrderBy {
	if value, ok := orderByValue[key]; ok {
		return OrderBy(value)
	}

	return OrderByAsc
}

func (o OrderBy) Name() string {
	if name, ok := orderByName[o]; ok {
		return name
	}

	return ""
}

func (o OrderBy) Proto() user.OrderBy {
	return *user.OrderBy(o).Enum()
}

// Gender - 性別
type Gender int32

const (
	GenderUnknown Gender = 0 // 未選択
	GenderMan     Gender = 1 // 男性
	GenderWoman   Gender = 2 // 女性
)

var (
	genderByName = map[Gender]string{
		0: "unknown",
		1: "man",
		2: "woman",
	}
	genderByValue = map[string]int32{
		"unknown": 0,
		"man":     1,
		"woman":   2,
	}
)

func NewGender(g user.Gender) Gender {
	switch g {
	case user.Gender_GENDER_UNKNOWN:
		return GenderUnknown
	case user.Gender_GENDER_MAN:
		return GenderMan
	case user.Gender_GENDER_WOMAN:
		return GenderWoman
	default:
		return GenderUnknown
	}
}

func NewGenderByValue(key string) Gender {
	if value, ok := genderByValue[key]; ok {
		return Gender(value)
	}

	return GenderUnknown
}

func (g Gender) Name() string {
	if name, ok := genderByName[g]; ok {
		return name
	}

	return "unknown"
}

func (g Gender) Proto() user.Gender {
	return *user.Gender(g).Enum()
}

// Role - ユーザー権限
type Role int32

const (
	RoleUser      Role = 0 // ユーザー (default)
	RoleAdmin     Role = 1 // 管理者
	RoleDeveloper Role = 2 // 開発者
	RoleOperator  Role = 3 // 運用者
)

var (
	roleByName = map[Role]string{
		0: "user",
		1: "admin",
		2: "developer",
		3: "operator",
	}
	roleByValue = map[string]int32{
		"user":      0,
		"admin":     1,
		"developer": 2,
		"operator":  3,
	}
)

func NewRole(r user.Role) Role {
	switch r {
	case user.Role_ROLE_USER:
		return RoleUser
	case user.Role_ROLE_ADMIN:
		return RoleAdmin
	case user.Role_ROLE_DEVELOPER:
		return RoleDeveloper
	case user.Role_ROLE_OPERATOR:
		return RoleOperator
	default:
		return RoleUser
	}
}

func NewRoleByValue(key string) Role {
	if value, ok := roleByValue[key]; ok {
		return Role(value)
	}

	return RoleUser
}

func (r Role) Name() string {
	if name, ok := roleByName[r]; ok {
		return name
	}

	return ""
}

func (r Role) Proto() user.Role {
	return *user.Role(r).Enum()
}

// BookshelfStatus - 読書ステータス
type BookshelfStatus int32

const (
	BookshelfStatusNone    BookshelfStatus = 0 // 不明なステータス
	BookshelfStatusRead    BookshelfStatus = 1 // 読み終えた本
	BookshelfStatusReading BookshelfStatus = 2 // 読んでいる本
	BookshelfStatusStacked BookshelfStatus = 3 // 積読本
	BookshelfStatusWant    BookshelfStatus = 4 // 読みたい本
	BookshelfStatusRelease BookshelfStatus = 5 // 手放したい本
)

var (
	bookshelfStatusName = map[BookshelfStatus]string{
		0: "none",
		1: "read",
		2: "reading",
		3: "stack",
		4: "want",
		5: "release",
	}
	bookshelfStatusValue = map[string]int32{
		"none":    0,
		"read":    1,
		"reading": 2,
		"stack":   3,
		"want":    4,
		"release": 5,
	}
)

func NewBookshelfStatus(s book.BookshelfStatus) BookshelfStatus {
	switch s {
	case book.BookshelfStatus_BOOKSHELF_STATUS_READ:
		return BookshelfStatusRead
	case book.BookshelfStatus_BOOKSHELF_STATUS_READING:
		return BookshelfStatusReading
	case book.BookshelfStatus_BOOKSHELF_STATUS_STACKED:
		return BookshelfStatusStacked
	case book.BookshelfStatus_BOOKSHELF_STATUS_WANT:
		return BookshelfStatusWant
	case book.BookshelfStatus_BOOKSHELF_STATUS_RELEASE:
		return BookshelfStatusRelease
	case book.BookshelfStatus_BOOKSHELF_STATUS_NONE:
		return BookshelfStatusNone
	default:
		return BookshelfStatusNone
	}
}

func NewBookshelfStatusByValue(key string) BookshelfStatus {
	if value, ok := bookshelfStatusValue[key]; ok {
		return BookshelfStatus(value)
	}

	return BookshelfStatusNone
}

func (s BookshelfStatus) Name() string {
	if name, ok := bookshelfStatusName[s]; ok {
		return name
	}

	return "none"
}

func (s BookshelfStatus) Proto() book.BookshelfStatus {
	return *book.BookshelfStatus(s).Enum()
}
