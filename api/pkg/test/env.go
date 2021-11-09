package test

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type testEnv struct {
	DBSocket          string `envconfig:"TEST_DB_SOCKET" default:"tcp"`
	DBHost            string `envconfig:"TEST_DB_HOST" default:"127.0.0.1"`
	DBPort            string `envconfig:"TEST_DB_PORT" default:"3326"`
	DBUsername        string `envconfig:"TEST_DB_USERNAME" default:"root"`
	DBPassword        string `envconfig:"TEST_DB_PASSWORD" default:"12345678"`
	DBUserDB          string `envconfig:"TEST_DB_USER_DB" default:"users"`
	DBBookDB          string `envconfig:"TEST_DB_BOOK_DB" default:"books"`
	DBInformationDB   string `envconfig:"TEST_DB_INFORMATION_DB" default:"informations"`
	GCPServiceKeyJSON string `envconfig:"GCP_SERVICE_KEY_JSON" default:"{}"`
	FAEmulatorHost    string `envconfig:"FIREBASE_AUTH_EMULATOR_HOST" default:"127.0.0.1:9099"`
	FBEmulatorHost    string `envconfig:"FIRESTORE_EMULATOR_HOST" default:"127.0.0.1:9090"`
}

func newTestEnv() (*testEnv, error) {
	env := &testEnv{}
	err := envconfig.Process("", env)
	if err != nil {
		return nil, fmt.Errorf("test: failed to load test environment: %w", err)
	}

	if os.Getenv("FIREBASE_AUTH_EMULATOR_HOST") == "" {
		os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "127.0.0.1:9099")
	}

	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		os.Setenv("FIRESTORE_EMULATOR_HOST", "127.0.0.1:9090")
	}

	return env, nil
}
