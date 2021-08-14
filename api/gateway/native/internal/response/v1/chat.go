package v1

type ChatRoomResponse struct {
	ID            string           `json:"id"`
	Users         []*ChatRoomUser  `json:"users"`
	LatestMessage *ChatRoomMessage `json:"latestMessage"`
	CreatedAt     string           `json:"createdAt"`
	UpdatedAt     string           `json:"updatedAt"`
}

type ChatRoomUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type ChatRoomMessage struct {
	UserID    string `json:"userId"`
	Text      string `json:"text"`
	Image     string `json:"image"`
	CreatedAt string `json:"createdAt"`
}

type ChatRoomListResponse struct {
	Rooms []*ChatRoomListRoom `json:"rooms"`
}

type ChatRoomListRoom struct {
	ID            string               `json:"id"`
	Users         []*ChatRoomListUser  `json:"users"`
	LatestMessage *ChatRoomListMessage `json:"latestMessage"`
	CreatedAt     string               `json:"createdAt"`
	UpdatedAt     string               `json:"updatedAt"`
}

type ChatRoomListUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type ChatRoomListMessage struct {
	UserID    string `json:"userId"`
	Text      string `json:"text"`
	Image     string `json:"image"`
	CreatedAt string `json:"createdAt"`
}

type ChatMessageResponse struct {
	Text      string           `json:"text"`
	Image     string           `json:"image"`
	CreatedAt string           `json:"createdAt"`
	User      *ChatMessageUser `json:"user"`
}

type ChatMessageUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
