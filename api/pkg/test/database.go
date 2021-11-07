package test

import (
	"context"
	"fmt"

	fs "cloud.google.com/go/firestore"
	fb "firebase.google.com/go/v4"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/firebase"
	"github.com/calmato/gran-book/api/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/pkg/firebase/firestore"
	"github.com/golang/mock/gomock"
)

var (
	userTables = []string{
		"users",
		"relationships",
	}
	bookTables = []string{
		"reviews",
		"authors_books",
		"authors",
		"bookshelves",
		"books",
	}
	informationTables = []string{
		"inquiries",
	}
)

func NewDBMock(ctrl *gomock.Controller) (*DBMocks, error) {
	env, err := newTestEnv()
	if err != nil {
		return nil, err
	}

	up := &database.Params{
		Socket:        env.DBSocket,
		Host:          env.DBHost,
		Port:          env.DBPort,
		Database:      env.DBUserDB,
		Username:      env.DBUsername,
		Password:      env.DBPassword,
		DisableLogger: true,
	}
	udb, err := database.NewClient(up)
	if err != nil {
		return nil, err
	}

	bp := &database.Params{
		Socket:        env.DBSocket,
		Host:          env.DBHost,
		Port:          env.DBPort,
		Database:      env.DBBookDB,
		Username:      env.DBUsername,
		Password:      env.DBPassword,
		DisableLogger: true,
	}
	bdb, err := database.NewClient(bp)
	if err != nil {
		return nil, err
	}

	ip := &database.Params{
		Socket:        env.DBSocket,
		Host:          env.DBHost,
		Port:          env.DBPort,
		Database:      env.DBInformationDB,
		Username:      env.DBUsername,
		Password:      env.DBPassword,
		DisableLogger: true,
	}
	idb, err := database.NewClient(ip)
	if err != nil {
		return nil, err
	}

	return &DBMocks{
		UserDB:        udb,
		BookDB:        bdb,
		InformationDB: idb,
	}, nil
}

func (m *DBMocks) Delete(cli *database.Client, tables ...string) error {
	for _, table := range tables {
		sql := fmt.Sprintf("DELETE FROM %s", table)
		if err := cli.DB.Exec(sql).Error; err != nil {
			return err
		}
	}

	return nil
}

func (m *DBMocks) DeleteAll() error {
	err := m.Delete(m.UserDB, userTables...)
	if err != nil {
		return err
	}

	err = m.Delete(m.BookDB, bookTables...)
	if err != nil {
		return err
	}

	err = m.Delete(m.InformationDB, informationTables...)
	if err != nil {
		return err
	}

	return nil
}

func NewFirebaseMock(ctx context.Context) (*FirebaseMocks, error) {
	// エミュレータの環境変数読み込み用
	_, err := newTestEnv()
	if err != nil {
		return nil, err
	}

	// Firebaseの設定
	app, err := firebase.InitializeApp(ctx, &fb.Config{ProjectID: "project-test"})
	if err != nil {
		return nil, err
	}

	// Firebase Authenticationの設定
	fa, err := authentication.NewClient(ctx, app.App)
	if err != nil {
		return nil, err
	}

	// Firestoreの設定
	fs, err := firestore.NewClient(ctx, app.App)
	if err != nil {
		return nil, err
	}

	// Cloud Storageの設定
	// gcs, err := storage.NewClient(ctx, app.App, env.GCPStorageBucketName)
	// if err != nil {
	// 	return nil, err
	// }

	return &FirebaseMocks{
		Auth:      fa,
		Firestore: fs,
	}, nil
}

func (m *FirebaseMocks) DeleteAll(ctx context.Context) error {
	refs, err := m.Firestore.Client.Collections(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("firestore collections get all: %w", err)
	}

	for i := range refs {
		m.DeleteCollection(ctx, refs[i])
	}

	return nil
}

func (m *FirebaseMocks) DeleteCollection(ctx context.Context, ref *fs.CollectionRef) error {
	docRefs, err := ref.DocumentRefs(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("firestore documents get all: %w", err)
	}

	for i := range docRefs {
		err := m.DeleteDoc(ctx, docRefs[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *FirebaseMocks) DeleteDoc(ctx context.Context, ref *fs.DocumentRef) error {
	collections, err := ref.Collections(ctx).GetAll()
	if err != nil {
		return err
	}

	// サブコレクションの削除
	if len(collections) > 0 {
		for i := range collections {
			refs, err := collections[i].DocumentRefs(ctx).GetAll()
			if err != nil {
				return fmt.Errorf("firestore documents get all: %w", err)
			}

			for i := range refs {
				m.DeleteDoc(ctx, refs[i])
			}
		}
	}

	_, err = ref.Delete(ctx)
	return err
}
