package input

// ListUser - ユーザー一覧のリクエスト
type ListUser struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}

// ListUserByUserIDs - ユーザー一覧のリクエスト
type ListUserByUserIDs struct {
	UserIDs []string `json:"userIds" validate:"unique,dive,required,max=36"`
}

// ListFollow - フォローユーザ一覧のリクエスト
type ListFollow struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}

// ListFollower - フォロワーユーザ一覧のリクエスト
type ListFollower struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}

// SearchUser - ユーザー検索のリクエスト
type SearchUser struct {
	Limit     int    `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int    `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
	Field     string `json:"field" validate:"required,oneof=username email"`
	Value     string `json:"value" validate:"required"`
}
