package domain

// ListQuery - 一覧取得時のクエリ用構造体
type ListQuery struct {
	Limit  int64
	Offset int64
	Order  *QueryOrder
}

// QueryOrder - ソート用の構造体
type QueryOrder struct {
	By        string
	Direction string
}
