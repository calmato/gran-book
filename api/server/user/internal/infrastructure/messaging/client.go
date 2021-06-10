package messaging

import (
	"log"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type Client struct {
	messaging *expo.PushClient
}

func NewMessagingClient() *Client {
	// Create a new Expo SDK client
	pc := expo.NewPushClient(nil)

	return &Client{
		messaging: pc,
	}
}

func (c *Client) getValidToken(instanceIDs []string) []expo.ExponentPushToken {
	tokens := []expo.ExponentPushToken{}

	for _, instanceID := range instanceIDs {
		// To check the token is valid
		token, err := expo.NewExponentPushToken(instanceID)
		if err == nil {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func (c *Client) checkResponse(res expo.PushResponse, err error) error {
	// Check errors
	if err != nil {
		return exception.ErrorInOtherAPI.New(err)
	}

	// TODO: エラーの出し方、後で考える
	// Validate responses
	if res.ValidateResponse() != nil {
		log.Println(res.PushMessage.To, "failed")
	}

	return nil
}
