package v1

// チャットルーム情報
type ChatRoomResponse struct {
	ID            string           `json:"id"`            // チャットルームID
	Users         []*ChatRoomUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string           `json:"createdAt"`     // 作成日時
	UpdatedAt     string           `json:"updatedAt"`     // 更新日時
}

type ChatRoomMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

type ChatRoomUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// チャットルーム一覧
type ChatRoomListResponse struct {
	Rooms []*ChatRoomListRoom `json:"rooms"` // チャットルーム一覧
}

type ChatRoomListRoom struct {
	ID            string               `json:"id"`            // チャットルームID
	Users         []*ChatRoomListUser  `json:"usersList"`     // 参加者一覧
	LatestMessage *ChatRoomListMessage `json:"latestMessage"` // 最新のメッセージ
	CreatedAt     string               `json:"createdAt"`     // 作成日時
	UpdatedAt     string               `json:"updatedAt"`     // 更新日時
}

type ChatRoomListMessage struct {
	UserID    string `json:"userId"`    // ユーザーID
	Text      string `json:"text"`      // テキストメッセージ
	Image     string `json:"image"`     // 添付画像URL
	CreatedAt string `json:"createdAt"` // 送信日時
}

type ChatRoomListUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}

// チャットメッセージ情報
type ChatMessageResponse struct {
	Text      string           `json:"text"`      // テキストメッセージ
	Image     string           `json:"image"`     // 添付画像URL
	User      *ChatMessageUser `json:"user"`      // 送信者
	CreatedAt string           `json:"createdAt"` // 送信日時
}

type ChatMessageUser struct {
	ID           string `json:"id"`           // ユーザーID
	Username     string `json:"username"`     // 表示名
	ThumbnailURL string `json:"thumbnailUrl"` // サムネイルURL
}
