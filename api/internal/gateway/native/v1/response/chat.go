package response

import "github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"

// チャットルーム情報
type ChatRoomResponse struct {
	*entity.ChatRoom
}

// チャットルーム一覧
type ChatRoomListResponse struct {
	Rooms entity.ChatRooms `json:"rooms"` // チャットルーム一覧
}

// チャットメッセージ情報
type ChatMessageResponse struct {
	*entity.ChatMessage
}
