package v1

// チャットルーム情報
type ChatRoomV1Response struct {
	Id            string                      `json:"id,omitempty"`            // チャットルームID
	Users         []*ChatRoomV1Response_User  `json:"usersList,omitempty"`     // 参加者一覧
	LatestMessage *ChatRoomV1Response_Message `json:"latestMessage,omitempty"` // 最新のメッセージ
	CreatedAt     string                      `json:"createdAt,omitempty"`     // 作成日時
	UpdatedAt     string                      `json:"updatedAt,omitempty"`     // 更新日時
}

type ChatRoomV1Response_Message struct {
	UserId    string `json:"userId,omitempty"`    // ユーザーID
	Text      string `json:"text,omitempty"`      // テキストメッセージ
	Image     string `json:"image,omitempty"`     // 添付画像URL
	CreatedAt string `json:"createdAt,omitempty"` // 送信日時
}

type ChatRoomV1Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// チャットルーム一覧
type ChatRoomListV1Response struct {
	Rooms []*ChatRoomListV1Response_Room `json:"rooms,omitempty"` // チャットルーム一覧
}

type ChatRoomListV1Response_Room struct {
	Id            string                          `json:"id,omitempty"`            // チャットルームID
	Users         []*ChatRoomListV1Response_User  `json:"usersList,omitempty"`     // 参加者一覧
	LatestMessage *ChatRoomListV1Response_Message `json:"latestMessage,omitempty"` // 最新のメッセージ
	CreatedAt     string                          `json:"createdAt,omitempty"`     // 作成日時
	UpdatedAt     string                          `json:"updatedAt,omitempty"`     // 更新日時
}

type ChatRoomListV1Response_Message struct {
	UserId    string `json:"userId,omitempty"`    // ユーザーID
	Text      string `json:"text,omitempty"`      // テキストメッセージ
	Image     string `json:"image,omitempty"`     // 添付画像URL
	CreatedAt string `json:"createdAt,omitempty"` // 送信日時
}

type ChatRoomListV1Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}

// チャットメッセージ情報
type ChatMessageV1Response struct {
	Text      string                      `json:"text,omitempty"`      // テキストメッセージ
	Image     string                      `json:"image,omitempty"`     // 添付画像URL
	User      *ChatMessageV1Response_User `json:"user,omitempty"`      // 送信者
	CreatedAt string                      `json:"createdAt,omitempty"` // 送信日時
}

type ChatMessageV1Response_User struct {
	Id           string `json:"id,omitempty"`           // ユーザーID
	Username     string `json:"username,omitempty"`     // 表示名
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"` // サムネイルURL
}
