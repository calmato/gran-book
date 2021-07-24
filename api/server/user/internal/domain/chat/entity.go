package chat

import "time"

// Room - チャットルーム エンティティ
type Room struct {
	ID            string    `firestore:"id"`
	UserIDs       []string  `firestore:"users"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
	LatestMessage *Message  `firestore:"latestMassage"`
	InstanceIDs   []string  `firestore:"-"`
}

// Message - チャットメッセージ エンティティ
type Message struct {
	ID        string    `firestore:"id"`
	Text      string    `firestore:"text"`
	Image     string    `firestore:"image"`
	UserID    string    `firestore:"userId"`
	Username  string    `firestore:"-"`
	CreatedAt time.Time `firestore:"createdAt"`
}
