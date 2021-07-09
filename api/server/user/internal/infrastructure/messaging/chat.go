package messaging

import (
	"fmt"

	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type chatMessaging struct {
	client *Client
}

func NewChatMessaging() chat.Messaging {
	mc := NewMessagingClient()

	return &chatMessaging{
		client: mc,
	}
}

func (m *chatMessaging) PushCreateRoom(cr *chat.Room) error {
	tokens := m.client.getValidToken(cr.InstanceIDs)
	title := "新規チャットルーム"
	body := "新しいチャットルームが作成されました."

	pm := &expo.PushMessage{
		To:         tokens,
		Title:      title,
		Body:       body,
		Sound:      "default",
		TTLSeconds: 300, // 5 min
		Priority:   expo.DefaultPriority,
	}

	res, err := m.client.messaging.Publish(pm)
	return m.client.checkResponse(res, err)
}

func (m *chatMessaging) PushNewMessage(cr *chat.Room, cm *chat.Message) error {
	tokens := m.client.getValidToken(cr.InstanceIDs)
	title := "新規チャットメッセージ"
	body := fmt.Sprintf("%s: %s", cm.Username, cm.Text)

	pm := &expo.PushMessage{
		To:         tokens,
		Title:      title,
		Body:       body,
		Sound:      "default",
		TTLSeconds: 300, // 5 min
		Priority:   expo.DefaultPriority,
	}

	res, err := m.client.messaging.Publish(pm)
	return m.client.checkResponse(res, err)
}
