package entity

const (
	LIST_LIMIT_DEFAULT  = "100" // 一覧取得上限
	LIST_OFFSET_DEFAULT = "0"   // 一覧取得開始位置
)

// OrderBy - ソート順
type OrderBy int32

const (
	ORDER_BY_ASC  OrderBy = 0 // 昇順
	ORDER_BY_DESC OrderBy = 1 // 降順
)
