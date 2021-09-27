package chat

import (
	"time"

	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/chat"
)

// Room - チャットルーム エンティティ
type Room struct {
	ID            string    `firestore:"id"`
	UserIDs       []string  `firestore:"users"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
	LatestMessage *Message  `firestore:"latestMessage"`
	InstanceIDs   []string  `firestore:"-"`
}

type Rooms []*Room

func (r *Room) Proto() *pb.Room {
	var m *pb.Message
	if r.LatestMessage != nil {
		m = r.LatestMessage.Proto()
	}

	return &pb.Room{
		Id:            r.ID,
		UserIds:       r.UserIDs,
		CreatedAt:     datetime.FormatTime(r.CreatedAt),
		UpdatedAt:     datetime.FormatTime(r.UpdatedAt),
		LatestMessage: m,
	}
}

func (rs Rooms) Proto() []*pb.Room {
	res := make([]*pb.Room, len(rs))
	for i := range rs {
		res[i] = rs[i].Proto()
	}
	return res
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

func (m *Message) Proto() *pb.Message {
	return &pb.Message{
		Id:        m.ID,
		UserId:    m.UserID,
		Text:      m.Text,
		Image:     m.Image,
		CreatedAt: datetime.FormatTime(m.CreatedAt),
	}
}
