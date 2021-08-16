package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client - DB操作用のクライアントの構造体
type Client struct {
	DB *gorm.DB
}

// NewClient - DBクライアントの構造体
func NewClient(socket, host, port, database, username, password string) (*Client, error) {
	con := getConfig(socket, host, port, database, username, password)

	// プライマリレプリカの作成
	db, err := getDBClient(con)
	if err != nil {
		return nil, err
	}

	c := &Client{
		DB: db,
	}

	return c, nil
}

func getDBClient(config string) (*gorm.DB, error) {
	opt := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	return gorm.Open(mysql.Open(config), opt)
}

func getConfig(socket, host, port, database, username, password string) string {
	switch socket {
	case "tcp":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
			database,
		)
	case "unix":
		return fmt.Sprintf(
			"%s:%s@unix(%s)/%s?charset=utf8mb4&parseTime=true",
			username,
			password,
			host,
			database,
		)
	default:
		return ""
	}
}
