package database

import (
	"fmt"
	"strconv"

	"github.com/calmato/gran-book/api/server/user/pkg/array"
	"gorm.io/gorm"
)

// ListQuery - 一覧取得時のクエリ用構造体
type ListQuery struct {
	Limit      int
	Offset     interface{}
	Order      *OrderQuery
	Conditions []*ConditionQuery
}

// OrderQuery - ソートクエリ用の構造体
type OrderQuery struct {
	Field   string
	OrderBy int
}

// ConditionQuery - 絞り込み用の構造体
type ConditionQuery struct {
	Field    string
	Operator string // ==, !=, >、>=, <=, <, IN, LIKE, BETWEEN
	Value    interface{}
}

// ソート順
const (
	OrderByAsc  int = iota // 昇順
	OrderByDesc            // 降順
)

// 一覧取得クエリの作成用
func (c *Client) GetListQuery(table string, db *gorm.DB, q *ListQuery) *gorm.DB {
	db = db.Table(table)

	if q == nil {
		return db
	}

	for _, con := range q.Conditions {
		db = setWhere(db, con)
	}

	db = setOrder(db, q.Order)
	db = setLimit(db, q.Limit)
	db = setOffset(db, q.Offset)

	return db
}

// 検索クエリの作成用
func (c *Client) GetListCount(table string, db *gorm.DB, q *ListQuery) (int, error) {
	var count int64

	if q != nil {
		for _, v := range q.Conditions {
			db = setWhere(db, v)
		}
	}

	err := db.Table(table).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func setOrder(db *gorm.DB, q *OrderQuery) *gorm.DB {
	if q == nil || q.Field == "" {
		return db
	}

	switch q.OrderBy {
	case OrderByAsc:
		str := fmt.Sprintf("%s asc", q.Field)
		return db.Order(str)
	case OrderByDesc:
		str := fmt.Sprintf("%s desc", q.Field)
		return db.Order(str)
	default:
		return db
	}
}

func setWhere(db *gorm.DB, c *ConditionQuery) *gorm.DB {
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
		q := fmt.Sprintf("%s IN (?)", c.Field)
		vals, _ := array.ConvertStrings(c.Value)
		return db.Where(q, vals)
	case "LIKE":
		q := fmt.Sprintf("%s LIKE ?", c.Field)
		n := fmt.Sprintf("%%%s%%", c.Value) // e.g.) あいうえお -> %あいうえお%
		return db.Where(q, n)
	case "BETWEEN":
		q := fmt.Sprintf("%s BETWEEN ? AND ?", c.Field)
		vals, err := array.ConvertStrings(c.Value)
		if err != nil {
			return db
		}

		return db.Where(q, vals[0], vals[1])
	default:
		return db
	}
}

func setLimit(db *gorm.DB, limit int) *gorm.DB {
	if limit > 0 {
		return db.Limit(limit)
	}

	return db
}

func setOffset(db *gorm.DB, offset interface{}) *gorm.DB {
	switch v := offset.(type) {
	case int:
		return db.Offset(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return db
		}

		return db.Offset(i)
	default:
		return db
	}
}
