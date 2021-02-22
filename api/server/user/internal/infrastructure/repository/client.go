package repository

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Client - DB操作用クライアントの構造体
type Client struct {
	db *gorm.DB
}

// NewDBClient - DBクライアントの生成
func NewDBClient(socket, host, port, database, username, password string) (*Client, error) {
	db, err := gorm.Open("mysql", getDBConfig(socket, host, port, database, username, password))
	if err != nil {
		return &Client{}, err
	}

	db.LogMode(true)

	return &Client{db}, nil
}

func getDBConfig(socket, host, port, database, username, password string) string {
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

func getOrder(o *domain.QueryOrder) string {
	if o.By == "" || o.Description == "" {
		return ""
	}

	switch strings.ToLower(o.Description) {
	case "asc":
		return fmt.Sprintf("%s asc", o.By)
	case "desc":
		return fmt.Sprintf("%s desc", o.By)
	default:
		return ""
	}
}
