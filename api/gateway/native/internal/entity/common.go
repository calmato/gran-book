package entity

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
