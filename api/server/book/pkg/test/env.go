package test

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type testEnv struct {
	DBSocket   string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost     string `envconfig:"TEST_DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"TEST_DB_PORT" default:"3326"`
	DBUsername string `envconfig:"TEST_DB_USERNAME" default:"root"`
	DBPassword string `envconfig:"TEST_DB_PASSWORD" default:"12345678"`
	DBUserDB   string `envconfig:"TEST_DB_USER_DB" default:"users"`
	DBBookDB   string `envconfig:"TEST_DB_BOOK_DB" default:"books"`
}

func newTestEnv() (*testEnv, error) {
	env := &testEnv{}
	if err := envconfig.Process("", env); err != nil {
		return nil, fmt.Errorf("test: failed to load test environment: %w", err)
	}

	return env, nil
}
