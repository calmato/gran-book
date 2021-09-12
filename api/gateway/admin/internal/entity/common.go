package entity

import (
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/user"
)

const (
	ListLimitDefault  = "100" // 一覧取得上限
	ListOffsetDefault = "0"   // 一覧取得開始位置
)

// OrderBy - ソート順
type OrderBy int32

const (
	OrderByAsc  OrderBy = 0 // 昇順
	OrderByDesc OrderBy = 1 // 降順
)

var (
	orderByName = map[OrderBy]string{
		0: "asc",
		1: "desc",
	}
	orderByValue = map[string]int32{
		"asc":  0,
		"desc": 1,
	}
)

func (o OrderBy) Name() string {
	if name, ok := orderByName[o]; ok {
		return name
	}

	return ""
}

func (o OrderBy) Value(key string) OrderBy {
	if value, ok := orderByValue[key]; ok {
		return OrderBy(value)
	}

	return OrderByAsc
}

func (o OrderBy) Proto() user.OrderBy {
	return *user.OrderBy(o).Enum()
}

// Role - ユーザー権限
type Role int32

const (
	RoleUser      Role = 0 // ユーザー (default)
	RoleAdmin     Role = 1 // 管理者
	RoleDeveloper Role = 2 // 開発者
	RoleOperator  Role = 3 // 運用者
)

var (
	roleByName = map[Role]string{
		0: "user",
		1: "admin",
		2: "developer",
		3: "operator",
	}
	roleByValue = map[string]int32{
		"user":      0,
		"admin":     1,
		"developer": 2,
		"operator":  3,
	}
)

func NewRole(r user.Role) Role {
	switch r {
	case user.Role_ROLE_USER:
		return RoleUser
	case user.Role_ROLE_ADMIN:
		return RoleAdmin
	case user.Role_ROLE_DEVELOPER:
		return RoleDeveloper
	case user.Role_ROLE_OPERATOR:
		return RoleOperator
	default:
		return RoleUser
	}
}

func (r Role) Name() string {
	if name, ok := roleByName[r]; ok {
		return name
	}

	return ""
}

func (r Role) Value(key string) Role {
	if value, ok := roleByValue[key]; ok {
		return Role(value)
	}

	return RoleUser
}

func (r Role) Proto() user.Role {
	return *user.Role(r).Enum()
}
