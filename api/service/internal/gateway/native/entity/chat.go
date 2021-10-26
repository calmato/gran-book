package entity

import (
	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
)

type ChatRoom struct {
	ID            string       `json:"id"`            // チャットルームID
	Users         ChatUsers    `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string       `json:"createdAt"`     // 作成日時
	UpdatedAt     string       `json:"updatedAt"`     // 更新日時
}

type ChatRooms []*ChatRoom

type ChatMessage struct {
	Text      string    `json:"text"`      // テキストメッセージ
	Image     string    `json:"image"`     // 添付画像URL
	UserID    string    `json:"userId"`    // ユーザーID
	User      *ChatUser `json:"user"`      // 送信者
	CreatedAt string    `json:"createdAt"` // 送信日時
}

type ChatUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

type ChatUsers []*ChatUser

func NewChatRoom(r *entity.ChatRoom, um map[string]*entity.User) *ChatRoom {
	var u *entity.User
	var m *entity.ChatMessage
	if r.LatestMessage != nil {
		m = entity.NewChatMessage(r.LatestMessage)
		u = um[m.UserId]
	}
	return &ChatRoom{
		ID:            r.Id,
		Users:         newChatUsers(r.UserIds, um),
		LatestMessage: NewChatMessage(m, u),
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func NewChatRooms(rs entity.ChatRooms, um map[string]*entity.User) ChatRooms {
	res := make(ChatRooms, len(rs))
	for i := range rs {
		res[i] = NewChatRoom(rs[i], um)
	}
	return res
}

func NewChatMessage(m *entity.ChatMessage, u *entity.User) *ChatMessage {
	if m == nil {
		return &ChatMessage{}
	}

	return &ChatMessage{
		Text:      m.Text,
		Image:     m.Image,
		UserID:    m.UserId,
		User:      newChatUser(u),
		CreatedAt: m.CreatedAt,
	}
}

func newChatUser(u *entity.User) *ChatUser {
	if u == nil {
		return &ChatUser{
			Username: UnknownUserName,
		}
	}

	return &ChatUser{
		ID:           u.Id,
		Username:     u.Username,
		ThumbnailURL: u.ThumbnailUrl,
	}
}

func newChatUsers(userIDs []string, um map[string]*entity.User) ChatUsers {
	res := make(ChatUsers, len(userIDs))
	for i, userID := range userIDs {
		res[i] = newChatUser(um[userID])
	}
	return res
}
