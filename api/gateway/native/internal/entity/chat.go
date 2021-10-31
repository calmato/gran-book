package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
)

type ChatRoom struct {
	*pb.ChatRoom
}

type ChatRooms []*ChatRoom

func NewChatRoom(cr *pb.ChatRoom) *ChatRoom {
	return &ChatRoom{cr}
}

func NewChatRooms(crs []*pb.ChatRoom) ChatRooms {
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
	*pb.ChatMessage
}

func NewChatMessage(cm *pb.ChatMessage) *ChatMessage {
	return &ChatMessage{cm}
}
