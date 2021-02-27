package domain

// ListQuery - 一覧取得時のクエリ用構造体
type ListQuery struct {
	Limit      int64
	Offset     int64
	Order      *QueryOrder
	Conditions []*QueryCondition
}

// QueryOrder - ソート用の構造体
type QueryOrder struct {
	By        string
	Direction string
}

// QueryCondition - 絞り込み用の構造体
type QueryCondition struct {
	Field    string
	Operator string // ==, !=, >、>=, <=, <, IN, LIKE, BETWEEN
	Value    interface{}
}
