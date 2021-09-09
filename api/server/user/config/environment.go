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
	DBSocket             string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost               string `envconfig:"DB_HOST" default:""`
	DBPort               string `envconfig:"DB_PORT" default:"3306"`
	DBUsername           string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword           string `envconfig:"DB_PASSWORD" default:""`
	DBDatabase           string `envconfig:"DB_DATABASE" default:""`
	GCPServiceKeyJSON    string `envconfig:"GCP_SERVICE_KEY_JSON" required:"true"`
	GCPStorageBucketName string `envconfig:"GCP_STORAGE_BUCKET_NAME" default:""`
}

// LoadEnvironment - 環境変数の取得
func loadEnvironment() (*Environment, error) {
	env := &Environment{}
	if err := envconfig.Process("", env); err != nil {
		return env, fmt.Errorf("config: failed to LoadEnvironment: %w", err)
	}

	return env, nil
}
