package repository

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/xerrors"
)

const (
	defaultLimit          = 100
	defaultOrderBy        = "id"
	defaultOrderDirection = "asc"
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

// getListQuery - SELECT SQL文作成用メソッド
func (c *Client) getListQuery(q *domain.ListQuery) (*gorm.DB, error) {
	db := c.db

	if q == nil {
		return db, nil
	}

	// WHERE句の追加
	db, err := setWhere(db, q.Conditions)
	if err != nil {
		return nil, err
	}

	// ORDER句の追加
	if q.Order == nil {
		q.Order = &domain.QueryOrder{
			By:        defaultOrderBy,
			Direction: defaultOrderDirection,
		}
	}

	db = setOrder(db, q.Order)
	db = setLimit(db, q.Limit)
	db = setOffset(db, q.Offset)

	return db, nil
}

func (c *Client) getListCount(q *domain.ListQuery, model interface{}) (int64, error) {
	var count int64
	db := c.db.Model(model)

	if q == nil {
		err := db.Count(&count).Error
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	// WHERE句の追加
	db, err := setWhere(db, q.Conditions)
	if err != nil {
		return 0, err
	}

	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func setWhere(db *gorm.DB, cs []*domain.QueryCondition) (*gorm.DB, error) {
	for _, c := range cs {
		if c == nil {
			err := xerrors.New("QueryCondition is nil")
			return nil, err
		}

		switch c.Operator {
		case "==":
			q := fmt.Sprintf("%s = ?", c.Field)
			db = db.Where(q, c.Value)
		case "!=":
			q := fmt.Sprintf("%s <> ?", c.Field)
			db = db.Where(q, c.Value)
		case ">", "<", ">=", "<=":
			q := fmt.Sprintf("%s %s ?", c.Field, c.Operator)
			db = db.Where(q, c.Value)
		case "IN":
			q := fmt.Sprintf("%s IN ?", c.Field)
			db = db.Where(q, c.Value)
		case "LIKE":
			q := fmt.Sprintf("%s LIKE ?", c.Field)
			n := fmt.Sprintf("%%%s%%", c.Value) // e.g.) あいうえお -> %あいうえお%
			db = db.Where(q, n)
		case "BETWEEN":
			q := fmt.Sprintf("%s BETWEEN ? AND ?", c.Field)
			m, n := convertConditionValues(c)
			db = db.Where(q, m, n)
		default:
			err := xerrors.New("Operator in QueryCondition is invalid word")
			return nil, err
		}
	}

	return db, nil
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
	return db.Limit(limit)
}

func setOffset(db *gorm.DB, offset int64) *gorm.DB {
	return db.Offset(offset)
}

// convertConditionValues - BETWEENのとき、valueが配列になっているため
func convertConditionValues(q *domain.QueryCondition) (interface{}, interface{}) {
	var m, n interface{}
	switch q.Value.(type) {
	case []int32:
		m = q.Value.([]int32)[0]
		n = q.Value.([]int32)[1]
	case []int64:
		m = q.Value.([]int64)[0]
		n = q.Value.([]int64)[1]
	case []string:
		m = q.Value.([]string)[0]
		n = q.Value.([]string)[1]
	}

	return m, n
}
