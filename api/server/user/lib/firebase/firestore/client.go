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

// Query - Where()メソッドのフィルタリング使用するクエリ構造体
type Query struct {
	Field    string
	Operator string // <、<=、==、>、>=
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

// Get - 単一のドキュメントの内容を取得
func (f *Firestore) Get(ctx context.Context, collection string, document string) (*firestore.DocumentSnapshot, error) {
	doc, err := f.Client.Collection(collection).Doc(document).Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// GetAll - コレクション内のすべてのドキュメントを取得
func (f *Firestore) GetAll(ctx context.Context, collection string) *firestore.DocumentIterator {
	return f.Client.Collection(collection).Documents(ctx)
}

// GetAllFirst - 初めから指定の件数分ドキュメントを取得
func (f *Firestore) GetAllFirst(
	ctx context.Context, collection string, orderBy string, sort string, length int,
) ([]*firestore.DocumentSnapshot, error) {
	if orderBy == "" {
		orderBy = "id"
	}

	orderBySort := firestore.Asc
	switch sort {
	case sortByAsc:
		orderBySort = firestore.Asc
	case sortByDesc:
		orderBySort = firestore.Desc
	}

	dsnap := f.Client.Collection(collection).OrderBy(orderBy, orderBySort).Limit(length).Documents(ctx)

	docs, err := dsnap.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// GetAllFromStartAt - 指定した箇所から指定の件数分ドキュメントを取得
func (f *Firestore) GetAllFromStartAt(
	ctx context.Context, collection string, orderBy string, sort string, document string, length int,
) ([]*firestore.DocumentSnapshot, error) {
	col := f.Client.Collection(collection)

	doc, err := col.Doc(document).Get(ctx)
	if err != nil {
		return nil, err
	}

	if orderBy == "" {
		orderBy = "id"
	}

	orderBySort := firestore.Asc
	switch sort {
	case sortByAsc:
		orderBySort = firestore.Asc
	case sortByDesc:
		orderBySort = firestore.Desc
	}

	dsnap := col.OrderBy(orderBy, orderBySort).StartAfter(doc.Data()[orderBy]).Limit(length).Documents(ctx)

	docs, err := dsnap.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// GetByQuery - where()を使用して、特定の条件を満たすすべてのドキュメントを取得
func (f *Firestore) GetByQuery(ctx context.Context, collection string, query *Query) *firestore.DocumentIterator {
	return f.Client.Collection(collection).Where(query.Field, query.Operator, query.Value).Documents(ctx)
}

// GetByQueries - 複数のwhere()メソッドをつなぎ合わせて、特定の条件を満たすすべてのドキュメントを取得
func (f *Firestore) GetByQueries(ctx context.Context, collection string, queries []*Query) *firestore.DocumentIterator {
	c := f.Client.Collection(collection).Query

	for _, q := range queries {
		c = c.Where(q.Field, q.Operator, q.Value)
	}

	return c.Documents(ctx)
}

// Search - 前方一致検索
func (f *Firestore) Search(
	ctx context.Context, collection string, orderBy string, sort string, query string, length int,
) ([]*firestore.DocumentSnapshot, error) {
	if orderBy == "" {
		orderBy = "id"
	}

	orderBySort := firestore.Asc
	switch sort {
	case sortByAsc:
		orderBySort = firestore.Asc
	case sortByDesc:
		orderBySort = firestore.Desc
	}

	dsnap := f.Client.Collection(collection).OrderBy(orderBy, orderBySort).
		StartAt(query).EndAt(query + "\uf8ff").
		Limit(length).Documents(ctx)

	docs, err := dsnap.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// SearchFromStartAt - ()指定した箇所から)前方一致検索
func (f *Firestore) SearchFromStartAt(
	ctx context.Context, collection string, orderBy string, sort string, query string, document string, length int,
) ([]*firestore.DocumentSnapshot, error) {
	col := f.Client.Collection(collection)

	doc, err := col.Doc(document).Get(ctx)
	if err != nil {
		return nil, err
	}

	if orderBy == "" {
		orderBy = "id"
	}

	orderBySort := firestore.Asc
	switch sort {
	case sortByAsc:
		orderBySort = firestore.Asc
	case sortByDesc:
		orderBySort = firestore.Desc
	}

	dsnap := f.Client.Collection(collection).OrderBy(orderBy, orderBySort).
		StartAt(query).EndAt(query + "\uf8ff").
		StartAfter(doc.Data()[orderBy]).
		Limit(length).Documents(ctx)

	docs, err := dsnap.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
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
