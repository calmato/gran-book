package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// Client - Firebaseの構造体
type Client struct {
	App *firebase.App
}

// InitializeApp - Cloud Firebase SDKの初期化
func InitializeApp(ctx context.Context, config *firebase.Config, opts ...option.ClientOption) (*Client, error) {
	app, err := firebase.NewApp(ctx, config, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{app}, nil
}
