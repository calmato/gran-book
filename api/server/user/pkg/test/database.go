package test

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/golang/mock/gomock"
)

var (
	userTables = []string{
		"users",
		"relationships",
	}
)

func NewDBMock(ctrl *gomock.Controller) (*DBMocks, error) {
	env, err := newTestEnv()
	if err != nil {
		return nil, err
	}

	udbp := &database.Params{
		Socket:        env.DBSocket,
		Host:          env.DBHost,
		Port:          env.DBPort,
		Database:      env.DBUserDB,
		Username:      env.DBUsername,
		Password:      env.DBPassword,
		DisableLogger: true,
	}
	udb, err := database.NewClient(udbp)
	if err != nil {
		return nil, err
	}

	return &DBMocks{
		UserDB: udb,
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
	// Clean User DB
	for _, table := range userTables {
		sql := fmt.Sprintf("DELETE FROM %s", table)
		if err := m.UserDB.DB.Exec(sql).Error; err != nil {
			return err
		}
	}

	return nil
}
