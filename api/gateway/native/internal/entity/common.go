package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
)

const (
	ListLimitDefault  = "100" // 一覧取得上限
	ListOffsetDefault = "0"   // 一覧取得開始位置
)

// OrderBy - ソート順
type OrderBy int32

const (
	OrderByAsc  OrderBy = 0 // 昇順
	OrderByDesc OrderBy = 1 // 降順
)

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

func NewGender(g pb.Gender) Gender {
	switch g {
	case pb.Gender_GENDER_UNKNOWN:
		return GenderUnknown
	case pb.Gender_GENDER_MAN:
		return GenderMan
	case pb.Gender_GENDER_WOMAN:
		return GenderWoman
	default:
		return GenderUnknown
	}
}

func (g Gender) Name() string {
	if name, ok := genderByName[g]; ok {
		return name
	}

	return "unknown"
}

func (g Gender) Value(key string) Gender {
	if value, ok := genderByValue[key]; ok {
		return Gender(value)
	}

	return GenderUnknown
}

func (g Gender) Proto() pb.Gender {
	return *pb.Gender(g).Enum()
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

func NewBookshelfStatus(s pb.BookshelfStatus) BookshelfStatus {
	switch s {
	case pb.BookshelfStatus_BOOKSHELF_STATUS_READ:
		return BookshelfStatusRead
	case pb.BookshelfStatus_BOOKSHELF_STATUS_READING:
		return BookshelfStatusReading
	case pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED:
		return BookshelfStatusStacked
	case pb.BookshelfStatus_BOOKSHELF_STATUS_WANT:
		return BookshelfStatusWant
	case pb.BookshelfStatus_BOOKSHELF_STATUS_RELEASE:
		return BookshelfStatusRelease
	case pb.BookshelfStatus_BOOKSHELF_STATUS_NONE:
		return BookshelfStatusNone
	default:
		return BookshelfStatusNone
	}
}

func (s BookshelfStatus) Name() string {
	if name, ok := bookshelfStatusName[s]; ok {
		return name
	}

	return "none"
}

func (s BookshelfStatus) Value(key string) BookshelfStatus {
	if value, ok := bookshelfStatusValue[key]; ok {
		return BookshelfStatus(value)
	}

	return BookshelfStatusNone
}
