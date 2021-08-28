package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
)

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

func (r Role) Proto() pb.Role {
	return *pb.Role(r).Enum()
}
