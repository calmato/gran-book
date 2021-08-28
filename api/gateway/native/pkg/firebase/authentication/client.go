package authentication

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// Auth - Authenticationの構造体
type Auth struct {
	Client *auth.Client
}

// NewClient - Firebase Authenticationに接続
func NewClient(ctx context.Context, app *firebase.App) (*Auth, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &Auth{client}, nil
}

// VerifyIDToken - IDトークンを確認して、デコードされたトークンからデバイスのuidを取得
func (a *Auth) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	t, err := a.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return t.UID, nil
}

// GetUserByUID - UIDによるユーザー情報の取得
func (a *Auth) GetUserByUID(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := a.Client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUIDByEmail - メールアドレスによるユーザーUIDの取得
func (a *Auth) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	u, err := a.Client.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// CreateUser - 新しいFirebase Authenticationユーザーを作成
func (a *Auth) CreateUser(ctx context.Context, uid string, email string, password string) (string, error) {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	u, err := a.Client.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// UpdateEmail - メールアドレスの変更
func (a *Auth) UpdateEmail(ctx context.Context, uid string, email string) error {
	params := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(false)

	_, err := a.Client.UpdateUser(ctx, uid, params)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePassword - Passwordを変更
func (a *Auth) UpdatePassword(ctx context.Context, uid string, password string) error {
	params := (&auth.UserToUpdate{}).
		Password(password)

	_, err := a.Client.UpdateUser(ctx, uid, params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateActivated - アカウントの状態を変更
func (a *Auth) UpdateActivated(ctx context.Context, uid string, disabled bool) error {
	params := (&auth.UserToUpdate{}).
		Disabled(disabled)

	_, err := a.Client.UpdateUser(ctx, uid, params)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser - 既存のユーザーをuidで削除
func (a *Auth) DeleteUser(ctx context.Context, uid string) error {
	return a.Client.DeleteUser(ctx, uid)
}
