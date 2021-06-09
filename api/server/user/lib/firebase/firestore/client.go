package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

// Firestore - Firestoreの構造体
type Firestore struct {
	Client *firestore.Client
}

// Params - Where()メソッドで使用するクエリ構造体
type Params struct {
	OrderBy string
	SortBy  string
	Offset  string
	Limit   int
}

// Query - Where()メソッドのフィルタリング使用するクエリ構造体
type Query struct {
	Field    string
	Operator string // <, <=, >, >=, ==, !=, array-contains, array-contains-any, in, not-int
	Value    interface{}
}

const (
	sortByAsc  = "asc"
	sortByDesc = "desc"
)

// NewClient - Firestoreに接続
func NewClient(ctx context.Context, app *firebase.App) (*Firestore, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Firestore{client}, nil
}

// Close - Firestoreとの接続を終了
func (f *Firestore) Close() error {
	return f.Client.Close()
}

// List - コレクション内のドキュメントを福数取得
func (f *Firestore) List(
	ctx context.Context, collection string, params *Params, queries []*Query,
) ([]*firestore.DocumentSnapshot, error) {
	c, err := f.getQuery(ctx, collection, params)
	if err != nil {
		return nil, err
	}

	// Wher句
	for _, q := range queries {
		c = c.Where(q.Field, q.Operator, q.Value)
	}

	docs, err := c.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// Search - 前方一致検索
func (f *Firestore) Search(
	ctx context.Context, collection string, params *Params, value string,
) ([]*firestore.DocumentSnapshot, error) {
	c, err := f.getQuery(ctx, collection, params)
	if err != nil {
		return nil, err
	}

	docs, err := c.StartAt(value).EndAt(value + "\uf8ff").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// Get - 単一のドキュメントの内容を取得
func (f *Firestore) Get(ctx context.Context, collection string, document string) (*firestore.DocumentSnapshot, error) {
	doc, err := f.Client.Collection(collection).Doc(document).Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Set - 単一のドキュメントを作成または上書き
func (f *Firestore) Set(ctx context.Context, collection string, document string, data interface{}) error {
	if _, err := f.Client.Collection(collection).Doc(document).Set(ctx, data); err != nil {
		return err
	}

	return nil
}

// DeleteDoc - ドキュメントを削除
func (f *Firestore) DeleteDoc(ctx context.Context, collection string, document string) error {
	if _, err := f.Client.Collection(collection).Doc(document).Delete(ctx); err != nil {
		return err
	}

	return nil
}

/*
 * Private Method
 */
func (f *Firestore) getQuery(ctx context.Context, collection string, params *Params) (firestore.Query, error) {
	// ソート
	if params.OrderBy == "" {
		params.OrderBy = "id"
	}

	orderBySort := firestore.Asc
	switch params.SortBy {
	case sortByAsc:
		orderBySort = firestore.Asc
	case sortByDesc:
		orderBySort = firestore.Desc
	}

	c := f.Client.Collection(collection).OrderBy(params.OrderBy, orderBySort)

	// 取得開始位置
	if params.Offset != "" {
		doc, err := f.Client.Collection(collection).Doc(params.Offset).Get(ctx)
		if err != nil {
			return firestore.Query{}, err
		}

		c = c.StartAfter(doc.Data()[params.OrderBy])
	}

	// 取得上限数
	if params.Limit != 0 {
		c = c.Limit(params.Limit)
	}

	return c, nil
}
