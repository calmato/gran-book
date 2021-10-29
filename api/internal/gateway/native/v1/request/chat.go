package request

// チャットルーム作成
type CreateChatRoomRequest struct {
	UserIDs []string `json:"userIdsList" binding:"required"` // 参加ユーザーID一覧
}

// チャットメッセージ送信
type CreateChatMessageRequest struct {
	Text string `json:"text" binding:"required"` // テキストメッセージ
}
