package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
)

type ChatRoom struct {
	*pb.Room
}

type ChatRooms []*ChatRoom

func NewChatRoom(cr *pb.Room) *ChatRoom {
	return &ChatRoom{cr}
}

func NewChatRooms(crs []*pb.Room) ChatRooms {
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
	*pb.Message
}

func NewChatMessage(cm *pb.Message) *ChatMessage {
	return &ChatMessage{cm}
}
