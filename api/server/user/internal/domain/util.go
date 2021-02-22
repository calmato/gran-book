package domain

// ListQuery - 一覧取得時のクエリ用構造体
type ListQuery struct {
	Limit  int64
	Offset int64
	Order  *QueryOrder
}

// SearchQuery - 検索時のクエリ用構造体
type SearchQuery struct {
	Limit         int64
	Offset        int64
	Order         *QueryOrder
	Fields        []*QueryField
	FieldOperator string // and,or
}

// QueryField - 詳細検索用の構造体
type QueryField struct {
	Field    string
	Operator string // ==,!=,<,<=,>,>=,in,like
	Value    interface{}
}

// QueryOrder - ソート用の構造体
type QueryOrder struct {
	By        string
	Direction string
}
