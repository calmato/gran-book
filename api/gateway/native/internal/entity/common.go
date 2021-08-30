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

	return ""
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
