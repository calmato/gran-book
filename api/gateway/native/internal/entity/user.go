package entity

// Gender - 性別
type Gender int32

const (
	Gender_GENDER_UNKNOWN Gender = 0 // 未選択
	Gender_GENDER_MAN     Gender = 1 // 男性
	Gender_GENDER_WOMAN   Gender = 2 // 女性
)

// Role - ユーザー権限
type Role int32

const (
	Role_ROLE_USER      Role = 0 // ユーザー (default)
	Role_ROLE_ADMIN     Role = 1 // 管理者
	Role_ROLE_DEVELOPER Role = 2 // 開発者
	Role_ROLE_OPERATOR  Role = 3 // 運用者
)
