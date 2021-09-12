package v1

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
)

// チャットルーム情報
type ChatRoomResponse struct {
	ID            string           `json:"id"`            // チャットルームID
	Users         []*ChatRoomUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string           `json:"createdAt"`     // 作成日時
	UpdatedAt     string           `json:"updatedAt"`     // 更新日時
}

func NewChatRoomResponse(cr *entity.ChatRoom, um map[string]*entity.User) *ChatRoomResponse {
	return &ChatRoomResponse{
		ID:            cr.Id,
		Users:         newChatRoomUsers(cr.UserIds, um),
		LatestMessage: newChatRoomMessage(cr.LatestMessage),
		CreatedAt:     cr.CreatedAt,
		UpdatedAt:     cr.UpdatedAt,
	}
}

type ChatRoomMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

func newChatRoomMessage(cm *chat.Message) *ChatRoomMessage {
	if cm == nil {
		return &ChatRoomMessage{}
	}

	return &ChatRoomMessage{
		UserID:    cm.UserId,
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: cm.CreatedAt,
	}
}

type ChatRoomUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatRoomUser(u *entity.User) *ChatRoomUser {
	if u == nil {
		return &ChatRoomUser{
			Username: "unknown",
		}
	}

	return &ChatRoomUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func newChatRoomUsers(userIDs []string, um map[string]*entity.User) []*ChatRoomUser {
	res := make([]*ChatRoomUser, len(userIDs))
	for i := range userIDs {
		u := um[userIDs[i]]
		res[i] = newChatRoomUser(u)
	}
	return res
}

// チャットルーム一覧
type ChatRoomListResponse struct {
	Rooms []*ChatRoomListRoom `json:"rooms"` // チャットルーム一覧
}

func NewChatRoomListResponse(crs entity.ChatRooms, um map[string]*entity.User) *ChatRoomListResponse {
	return &ChatRoomListResponse{
		Rooms: newChatRoomListRooms(crs, um),
	}
}

type ChatRoomListRoom struct {
	ID            string               `json:"id"`            // チャットルームID
	Users         []*ChatRoomListUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomListMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string               `json:"createdAt"`     // 作成日時
	UpdatedAt     string               `json:"updatedAt"`     // 更新日時
}

func newChatRoomListRoom(cr *entity.ChatRoom, um map[string]*entity.User) *ChatRoomListRoom {
	return &ChatRoomListRoom{
		ID:            cr.Id,
		Users:         newChatRoomListUsers(cr.UserIds, um),
		LatestMessage: newChatRoomListMessage(cr.LatestMessage),
		CreatedAt:     cr.CreatedAt,
		UpdatedAt:     cr.UpdatedAt,
	}
}

func newChatRoomListRooms(crs entity.ChatRooms, um map[string]*entity.User) []*ChatRoomListRoom {
	res := make([]*ChatRoomListRoom, len(crs))
	for i := range crs {
		res[i] = newChatRoomListRoom(crs[i], um)
	}
	return res
}

type ChatRoomListMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

func newChatRoomListMessage(cm *chat.Message) *ChatRoomListMessage {
	if cm == nil {
		return &ChatRoomListMessage{}
	}

	return &ChatRoomListMessage{
		UserID:    cm.UserId,
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: cm.CreatedAt,
	}
}

type ChatRoomListUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatRoomListUser(u *entity.User) *ChatRoomListUser {
	if u == nil {
		return &ChatRoomListUser{
			Username: "unknown",
		}
	}

	return &ChatRoomListUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func newChatRoomListUsers(userIDs []string, um map[string]*entity.User) []*ChatRoomListUser {
	res := make([]*ChatRoomListUser, len(userIDs))
	for i := range userIDs {
		u := um[userIDs[i]]
		res[i] = newChatRoomListUser(u)
	}
	return res
}

// チャットメッセージ情報
type ChatMessageResponse struct {
	Text      string           `json:"text"`      // テキストメッセージ
	Image     string           `json:"image"`     // 添付画像URL
	User      *ChatMessageUser `json:"user"`      // 送信者
	CreatedAt string           `json:"createdAt"` // 送信日時
}

func NewChatMessageResponse(cm *entity.ChatMessage, a *entity.Auth) *ChatMessageResponse {
	return &ChatMessageResponse{
		Text:      cm.Text,
		Image:     cm.Image,
		User:      newChatMessageUser(a),
		CreatedAt: cm.CreatedAt,
	}
}

type ChatMessageUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatMessageUser(a *entity.Auth) *ChatMessageUser {
	return &ChatMessageUser{
		ID:           a.Id,
		Username:     a.Username,
		ThumbnailURL: a.ThumbnailUrl,
	}
}
