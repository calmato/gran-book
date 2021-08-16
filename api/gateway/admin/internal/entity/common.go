package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
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

func (o OrderBy) Name() string {
	if name, ok := orderByName[o]; ok {
		return name
	}

	return ""
}

func (o OrderBy) Value(key string) OrderBy {
	if value, ok := orderByValue[key]; ok {
		return OrderBy(value)
	}

	return OrderByAsc
}

func (o OrderBy) Proto() pb.OrderBy {
	return *pb.OrderBy(o).Enum()
}
