package entity

const (
	LIST_LIMIT_DEFAULT  = "100" // 一覧取得上限
	LIST_OFFSET_DEFAULT = "0"   // 一覧取得開始位置
)

// OrderBy - ソート順
type OrderBy int32

const (
	ORDER_BY_ASC  OrderBy = 0 // 昇順
	ORDER_BY_DESC OrderBy = 1 // 降順
)

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
