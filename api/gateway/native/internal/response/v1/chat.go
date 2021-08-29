package v1

// チャットルーム情報
type ChatRoomResponse struct {
	ID            string                    `json:"id"`            // チャットルームID
	Users         []*ChatRoomResponse_User  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomResponse_Message `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string                    `json:"createdAt"`     // 作成日時
	UpdatedAt     string                    `json:"updatedAt"`     // 更新日時
}

type ChatRoomResponse_Message struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

type ChatRoomResponse_User struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// チャットルーム一覧
type ChatRoomListResponse struct {
	Rooms []*ChatRoomListResponse_Room `json:"rooms"` // チャットルーム一覧
}

type ChatRoomListResponse_Room struct {
	ID            string                        `json:"id"`            // チャットルームID
	Users         []*ChatRoomListResponse_User  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomListResponse_Message `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string                        `json:"createdAt"`     // 作成日時
	UpdatedAt     string                        `json:"updatedAt"`     // 更新日時
}

type ChatRoomListResponse_Message struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

type ChatRoomListResponse_User struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl"` // サムネイルURL
}

// チャットメッセージ情報
type ChatMessageResponse struct {
	Text      string                    `json:"text"`      // テキストメッセージ
	Image     string                    `json:"image"`     // 添付画像URL
	User      *ChatMessageResponse_User `json:"user"`      // 送信者
	CreatedAt string                    `json:"createdAt"` // 送信日時
}

type ChatMessageResponse_User struct {
	Id           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl"` // サムネイルURL
}
