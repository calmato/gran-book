package chat

type Messaging interface {
	PushNewMessage(cr *Room, cm *Message) error
}
