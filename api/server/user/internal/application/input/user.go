package input

// ListFollow - フォローユーザ一覧のリクエスト
type ListFollow struct {
	Limit     int64  `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int64  `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}

// ListFollower - フォロワーユーザ一覧のリクエスト
type ListFollower struct {
	Limit     int64  `json:"limit" validate:"gte=0,lte=1000"`
	Offset    int64  `json:"offset" validate:"gte=0"`
	By        string `json:"by" validate:"omitempty,oneof=id username email role"`
	Direction string `json:"direction" validate:"omitempty,oneof=asc desc"`
}
