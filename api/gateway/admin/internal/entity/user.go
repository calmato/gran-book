package entity

// Role - ユーザー権限
type Role int32

const (
	RoleUser      Role = 0 // ユーザー (default)
	RoleAdmin     Role = 1 // 管理者
	RoleDeveloper Role = 2 // 開発者
	RoleOperator  Role = 3 // 運用者
)
