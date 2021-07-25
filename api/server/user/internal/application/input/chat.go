package input

// CreateRoom - チャットルーム作成のリクエスト
type CreateRoom struct {
	UserIDs []string `json:"userIds" validate:"unique,dive,required,max=36"`
}

// CreateTextMessage - チャットメッセージ(テキスト)作成のリクエスト
type CreateTextMessage struct {
	Text string `json:"text" validate:"required,max=1000"`
}

// CreateImageMessage - チャットメッセージ(テキスト)作成のリクエスト
type CreateImageMessage struct {
	Image []byte `json:"image" validate:"required"`
}
