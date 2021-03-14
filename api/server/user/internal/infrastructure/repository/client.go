package repository

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/lib/array"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client - DB操作用クライアントの構造体
type Client struct {
	db *gorm.DB
}

// NewDBClient - DBクライアントの生成
func NewDBClient(socket, host, port, database, username, password string) (*Client, error) {
	dsn := getDBConfig(socket, host, port, database, username, password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return &Client{}, err
	}

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

// getListQuery - SELECT SQL文作成用メソッド
func (c *Client) getListQuery(db *gorm.DB, q *domain.ListQuery) *gorm.DB {
	if q == nil {
		return db
	}

	// WHERE句の追加
	for _, c := range q.Conditions {
		db = setWhere(db, c)
	}

	db = setOrder(db, q.Order)
	db = setLimit(db, q.Limit)
	db = setOffset(db, q.Offset)

	return db
}

func (c *Client) getListCount(db *gorm.DB, q *domain.ListQuery) (int64, error) {
	var count int64

	if q != nil {
		// WHERE句の追加
		for _, c := range q.Conditions {
			db = setWhere(db, c)
		}
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return count, nil
}

func setWhere(db *gorm.DB, c *domain.QueryCondition) *gorm.DB {
	if c == nil {
		return db
	}

	switch c.Operator {
	case "==":
		q := fmt.Sprintf("%s = ?", c.Field)
		return db.Where(q, c.Value)
	case "!=":
		q := fmt.Sprintf("%s <> ?", c.Field)
		return db.Where(q, c.Value)
	case ">", "<", ">=", "<=":
		q := fmt.Sprintf("%s %s ?", c.Field, c.Operator)
		return db.Where(q, c.Value)
	case "IN":
		q := fmt.Sprintf("%s IN ?", c.Field)
		vals, _ := array.ConvertStrings(c)
		return db.Where(q, strings.Join(vals, ", "))
	case "LIKE":
		q := fmt.Sprintf("%s LIKE ?", c.Field)
		n := fmt.Sprintf("%%%s%%", c.Value) // e.g.) あいうえお -> %あいうえお%
		return db.Where(q, n)
	case "BETWEEN":
		q := fmt.Sprintf("%s BETWEEN ? AND ?", c.Field)
		vals, _ := array.ConvertStrings(c)
		return db.Where(q, vals[0], vals[1])
	default:
		return db
	}
}

func setOrder(db *gorm.DB, o *domain.QueryOrder) *gorm.DB {
	if o == nil || o.By == "" {
		return db
	}

	switch strings.ToLower(o.Direction) {
	case "asc":
		q := fmt.Sprintf("%s asc", o.By)
		return db.Order(q)
	case "desc":
		q := fmt.Sprintf("%s desc", o.By)
		return db.Order(q)
	default:
		return db
	}
}

func setLimit(db *gorm.DB, limit int64) *gorm.DB {
	if limit > 0 {
		return db.Limit(int(limit))
	}

	return db
}

func setOffset(db *gorm.DB, offset int64) *gorm.DB {
	return db.Offset(int(offset))
}
