package output

// ListQuery - 一覧検索クエリのレスポンス
type ListQuery struct {
	Limit  int
	Offset int
	Total  int
	Order  *QueryOrder
}

// QueryOrder - 検索クエリソートのレスポンス
type QueryOrder struct {
	By        string
	Direction string
}
