package chat

import "time"

// Room - チャットルーム エンティティ
type Room struct {
	ID            string    `firestore:"id"`
	UserIDs       []string  `firestore:"users"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
	LatestMassage *Message  `firestore:"-"`
}

// Message - チャットメッセージ エンティティ
type Message struct {
	ID        string    `firestore:"id"`
	Text      string    `firestore:"text"`
	Image     string    `firestore:"image"`
	UserID    string    `firestore:"userId"`
	CreatedAt time.Time `firestore:"createdAt"`
}
