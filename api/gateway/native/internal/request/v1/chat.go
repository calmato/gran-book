package v1

type CreateChatRoomRequest struct {
	UserIDs []string `json:"users"`
}

type CreateChatMessageRequest struct {
	Text string `json:"text"`
}
