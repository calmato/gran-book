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
	for _, condition := range q.Conditions {
		c, err := getConditionQuery(condition)
		if err != nil {
			return nil, err
		}

		switch condition.Operator {
		case "BETWEEN":
			m, n := convertConditionValues(condition)
			db = db.Where(c, m, n)
		default:
			db = db.Where(c, condition.Value)
		}
	}

	// ORDER句の追加
	if q.Order == nil {
		q.Order = &domain.QueryOrder{
			By:        defaultOrderBy,
			Direction: defaultOrderDirection,
		}
	}

	db = db.Order(getOrder(q.Order))

	// LIMIT句の追加
	if q.Limit == 0 {
		db = db.Limit(q.Limit)
	} else {
		db = db.Limit(defaultLimit)
	}

	// OFFSET句の追加
	if q.Offset == 0 {
		db = db.Offset(q.Offset)
	}

	return db, nil
}

// getConditionQuery - WHERE句の作成用
func getConditionQuery(q *domain.QueryCondition) (string, error) {
	if q == nil {
		err := xerrors.New("QueryCondition is nil")
		return "", err
	}

	switch q.Operator {
	case "==", "!=", ">", "<", ">=", "<=":
		return fmt.Sprintf("%s %s ?", q.Field, q.Operator), nil
	case "IN", "LIKE":
		return fmt.Sprintf("%s %s ?", q.Field, q.Operator), nil
	case "BETWEEN":
		return fmt.Sprintf("%s BETWEEN ? AND ?", q.Field), nil
	default:
		err := xerrors.New("Operator in QueryCondition is invalid word")
		return "", err
	}
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

// getOrder - ORDER句の作成用
func getOrder(o *domain.QueryOrder) string {
	if o == nil {
		return ""
	}

	if o.By == "" || o.Direction == "" {
		return ""
	}

	switch strings.ToLower(o.Direction) {
	case "asc":
		return fmt.Sprintf("%s asc", o.By)
	case "desc":
		return fmt.Sprintf("%s desc", o.By)
	default:
		return ""
	}
}
