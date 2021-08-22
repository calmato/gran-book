package test

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/golang/mock/gomock"
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

func (m *DBMocks) DeleteAll(cli *database.Client, tables ...string) error {
	for _, table := range tables {
		sql := fmt.Sprintf("DELETE FROM %s", table)
		if err := cli.DB.Exec(sql).Error; err != nil {
			return err
		}
	}

	return nil
}
