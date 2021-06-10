package messaging

import (
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

func (m *chatMessaging) PushNewMessage(instanceIDs []string, cr *chat.Room, cm *chat.Message) error {
	tokens := m.client.getToken(instanceIDs)

	// TODO: チャットルーム名カラムを追加 -> titleにチャットルーム名が入るようにする
	title := "テストメッセージ"

	body := ""
	if cm.Text != "" {
		body = cm.Text
	} else if cm.Image != "" {
		body = "画像が送信されました."
	} else {
		body = "新規メッセージがあります."
	}

	// TODO: いったん1対1でのチャットのみを想定したデータ構造にしておく -> 後日リファクタ
	res, err := m.client.messaging.Publish(
		&expo.PushMessage{
			To:       tokens,
			Title:    title,
			Body:     body,
			Sound:    "default",
			Ttl:      300, // 5 min
			Priority: expo.DefaultPriority,
		},
	)

	return m.client.checkResponse(res, err)
}
