package chat

type Messaging interface {
	PushNewMessage(instanceIDs []string, cr *Room, cm *Message) error
}
