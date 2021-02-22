package output

// ListQuery - 一覧検索クエリのレスポンス
type ListQuery struct {
	Limit  int64
	Offset int64
	Total  int64
	Order  QueryOrder
}

// QueryOrder - 検索クエリソートのレスポンス
type QueryOrder struct {
	By        string
	Direction string
}
