package entity

// Gender - 性別
type Gender int32

const (
	GenderUnknown Gender = 0 // 未選択
	GenderMan     Gender = 1 // 男性
	GenderWoman   Gender = 2 // 女性
)

// Role - ユーザー権限
type Role int32

const (
	RoleUser      Role = 0 // ユーザー (default)
	RoleAdmin     Role = 1 // 管理者
	RoleDeveloper Role = 2 // 開発者
	RoleOperator  Role = 3 // 運用者
)
