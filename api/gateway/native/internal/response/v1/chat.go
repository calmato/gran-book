package v1

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
)

// チャットルーム情報
type ChatRoomResponse struct {
	ID            string           `json:"id"`            // チャットルームID
	Users         []*chatRoomUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *chatRoomMessage `json:"latestMessage"` // 最新のメッセージ
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

type chatRoomMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

func newChatRoomMessage(cm *chat.Message) *chatRoomMessage {
	if cm == nil {
		return &chatRoomMessage{}
	}

	return &chatRoomMessage{
		UserID:    cm.UserId,
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: cm.CreatedAt,
	}
}

type chatRoomUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatRoomUser(u *entity.User) *chatRoomUser {
	if u == nil {
		return &chatRoomUser{
			Username: "unknown",
		}
	}

	return &chatRoomUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func newChatRoomUsers(userIDs []string, um map[string]*entity.User) []*chatRoomUser {
	res := make([]*chatRoomUser, len(userIDs))
	for i := range userIDs {
		u := um[userIDs[i]]
		res[i] = newChatRoomUser(u)
	}
	return res
}

// チャットルーム一覧
type ChatRoomListResponse struct {
	Rooms []*chatRoomListRoom `json:"rooms"` // チャットルーム一覧
}

func NewChatRoomListResponse(crs entity.ChatRooms, um map[string]*entity.User) *ChatRoomListResponse {
	return &ChatRoomListResponse{
		Rooms: newChatRoomListRooms(crs, um),
	}
}

type chatRoomListRoom struct {
	ID            string               `json:"id"`            // チャットルームID
	Users         []*chatRoomListUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *chatRoomListMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string               `json:"createdAt"`     // 作成日時
	UpdatedAt     string               `json:"updatedAt"`     // 更新日時
}

func newChatRoomListRoom(cr *entity.ChatRoom, um map[string]*entity.User) *chatRoomListRoom {
	return &chatRoomListRoom{
		ID:            cr.Id,
		Users:         newChatRoomListUsers(cr.UserIds, um),
		LatestMessage: newChatRoomListMessage(cr.LatestMessage),
		CreatedAt:     cr.CreatedAt,
		UpdatedAt:     cr.UpdatedAt,
	}
}

func newChatRoomListRooms(crs entity.ChatRooms, um map[string]*entity.User) []*chatRoomListRoom {
	res := make([]*chatRoomListRoom, len(crs))
	for i := range crs {
		res[i] = newChatRoomListRoom(crs[i], um)
	}
	return res
}

type chatRoomListMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

func newChatRoomListMessage(cm *chat.Message) *chatRoomListMessage {
	if cm == nil {
		return &chatRoomListMessage{}
	}

	return &chatRoomListMessage{
		UserID:    cm.UserId,
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: cm.CreatedAt,
	}
}

type chatRoomListUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatRoomListUser(u *entity.User) *chatRoomListUser {
	if u == nil {
		return &chatRoomListUser{
			Username: "unknown",
		}
	}

	return &chatRoomListUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func newChatRoomListUsers(userIDs []string, um map[string]*entity.User) []*chatRoomListUser {
	res := make([]*chatRoomListUser, len(userIDs))
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
	User      *chatMessageUser `json:"user"`      // 送信者
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

type chatMessageUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

func newChatMessageUser(a *entity.Auth) *chatMessageUser {
	return &chatMessageUser{
		ID:           a.Id,
		Username:     a.Username,
		ThumbnailURL: a.ThumbnailUrl,
	}
}
