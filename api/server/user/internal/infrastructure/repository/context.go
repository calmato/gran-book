package repository

import "strings"

const (
	chatRoomCollection    = "rooms"
	chatMessageCollection = "messages"
)

func getChatRoomCollection() string {
	return chatRoomCollection
}

func getChatMessageCollection(roomID string) string {
	return strings.Join([]string{chatRoomCollection, roomID, chatMessageCollection}, "/")
}
