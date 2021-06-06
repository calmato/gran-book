package input

// CreateRoom - チャットルーム作成のリクエスト
type CreateRoom struct {
	UserIDs []string `json:"userIds" validate:"unique,dive,required,max=36"`
}
