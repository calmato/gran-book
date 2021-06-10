package chat

type Messaging interface {
	PushCreateRoom(cr *Room) error
	PushNewMessage(cr *Room, cm *Message) error
}
