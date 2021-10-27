package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client - DB操作用のクライアント構造体
type Client struct {
	DB *gorm.DB
}

type Params struct {
	Socket        string
	Host          string
	Port          string
	Database      string
	Username      string
	Password      string
	DisableLogger bool
}

// NewClient - DBクライアントの構造体
func NewClient(params *Params) (*Client, error) {
	con := getConfig(params)

	// プライマリレプリカの作成
	db, err := getDBClient(con, params)
	if err != nil {
		return nil, err
	}

	c := &Client{
		DB: db,
	}

	return c, nil
}

func (c *Client) Begin(opts ...*sql.TxOptions) (*gorm.DB, error) {
	tx := c.DB.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) Close(tx *gorm.DB) func() {
	return func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}
}

func getDBClient(config string, params *Params) (*gorm.DB, error) {
	opt := &gorm.Config{}

	if !params.DisableLogger {
		opt.Logger = logger.Default.LogMode(logger.Info)
	}

	return gorm.Open(mysql.Open(config), opt)
}

func getConfig(params *Params) string {
	switch params.Socket {
	case "tcp":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			params.Username,
			params.Password,
			params.Host,
			params.Port,
			params.Database,
		)
	case "unix":
		return fmt.Sprintf(
			"%s:%s@unix(%s)/%s?charset=utf8mb4&parseTime=true",
			params.Username,
			params.Password,
			params.Host,
			params.Database,
		)
	default:
		return ""
	}
}
