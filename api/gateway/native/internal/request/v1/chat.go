package v1

// チャットルーム作成
type CreateChatRoomRequest struct {
	UserIds []string `json:"userIds"` // 参加ユーザーID一覧
}

// チャットメッセージ送信
type CreateChatMessageRequest struct {
	Text string `json:"text"` // テキストメッセージ
}
