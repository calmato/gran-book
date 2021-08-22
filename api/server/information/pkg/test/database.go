package test

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/information/pkg/database"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
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

	idbp := &database.Params{
		Socket:        env.DBSocket,
		Host:          env.DBHost,
		Port:          env.DBPort,
		Database:      env.DBInfomationDB,
		Username:      env.DBUsername,
		Password:      env.DBPassword,
		DisableLogger: true,
	}
	idb, err := database.NewClient(idbp)
	if err != nil {
		return nil, err
	}

	return &DBMocks{
		UserDB:        udb,
		InformationDB: idb,
	}, nil
}

func (m *DBMocks) CreateUser() (string, error) {
	userID := uuid.New().String()

	sql := "INSERT INTO users (id, username, created_at, updated_at) VALUES (?, ?, ?, ?)"
	err := m.UserDB.DB.Exec(sql, userID, "test-user", TimeMock, TimeMock).Error
	if err != nil {
		return "", nil
	}

	return userID, nil
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
