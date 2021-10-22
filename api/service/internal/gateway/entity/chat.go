package entity

import (
	"github.com/calmato/gran-book/api/service/proto/chat"
)

type ChatRoom struct {
	*chat.Room
}

type ChatRooms []*ChatRoom

func NewChatRoom(cr *chat.Room) *ChatRoom {
	return &ChatRoom{cr}
}

func NewChatRooms(crs []*chat.Room) ChatRooms {
	res := make(ChatRooms, len(crs))
	for i := range crs {
		res[i] = NewChatRoom(crs[i])
	}
	return res
}

func (crs ChatRooms) UserIDs() []string {
	userIDs := []string{}
	users := map[string]bool{}
	for _, cr := range crs {
		for _, userID := range cr.UserIds {
			if _, ok := users[userID]; ok {
				continue
			}

			users[userID] = true
			userIDs = append(userIDs, userID)
		}
	}
	return userIDs
}

type ChatMessage struct {
	*chat.Message
}

func NewChatMessage(cm *chat.Message) *ChatMessage {
	return &ChatMessage{cm}
}
