package test

import (
	"fmt"

	"github.com/calmato/gran-book/api/pkg/database"
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

	return &DBMocks{
		UserDB: udb,
		BookDB: bdb,
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

	return nil
}
