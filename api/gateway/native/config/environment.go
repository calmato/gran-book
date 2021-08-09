package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Environment - システム内で利用する環境変数の構造体
type Environment struct {
	Port                 string `envconfig:"PORT" default:"8080"`
	MetricsPort          string `envconfig:"METRICS_PORT" default:"9090"`
	LogPath              string `envconfig:"LOG_PATH" default:""`
	LogLevel             string `envconfig:"LOG_LEVEL" default:"info"`
	AuthServiceURL       string `envconfig:"AUTH_SERVICE_URL" default:"user_api:8080"`
	UserServiceURL       string `envconfig:"USER_SERVICE_URL" default:"user_api:8080"`
	GCPServiceKeyJSON    string `envconfig:"GCP_SERVICE_KEY_JSON" required:"true"`
	GCPStorageBucketName string `envconfig:"GCP_STORAGE_BUCKET_NAME" default:""`
}

// loadEnvironment - 環境変数の取得
func loadEnvironment() (*Environment, error) {
	env := &Environment{}
	if err := envconfig.Process("", env); err != nil {
		return env, fmt.Errorf("failed to LoadEnvironment: %+v", err)
	}

	return env, nil
}
