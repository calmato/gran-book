package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/calmato/gran-book/api/server/notification/internal/domain/auth"
	"github.com/calmato/gran-book/api/server/notification/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/notification/lib/metadata"
	"golang.org/x/xerrors"
)

type authRepository struct{}

type firebaseToken struct {
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	UID      string                 `json:"uid,omitempty"`
	Claims   map[string]interface{} `json:"-"`
}

// NewAuthRepository - AuthRepositoryの生成
func NewAuthRepository() auth.Repository {
	return &authRepository{}
}

func (r *authRepository) Authentication(ctx context.Context) (string, error) {
	t, err := getToken(ctx)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	fbToken, err := decodeToken(t)
	if err != nil {
		return "", exception.Unauthorized.New(err)
	}

	return fbToken.Subject, nil
}

func getToken(ctx context.Context) (string, error) {
	authorization, err := metadata.Get(ctx, "Authorization")
	if err != nil {
		return "", err
	}

	if authorization == "" {
		return "", xerrors.New("Authorization header is not contain.")
	}

	t := strings.Replace(authorization, "Bearer ", "", 1)
	return t, nil
}

func decodeToken(token string) (*firebaseToken, error) {
	s := strings.Split(token, ".")
	if len(s) != 3 {
		return nil, xerrors.New("Authorization header is invalid.")
	}

	data, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return nil, err
	}

	fbToken := &firebaseToken{}

	err = json.Unmarshal(data, fbToken)
	if err != nil {
		return nil, err
	}

	err = verifyToken(fbToken)
	if err != nil {
		return nil, err
	}

	return fbToken, nil
}

func verifyToken(t *firebaseToken) error {
	now := time.Now().Unix()

	verifyTokenMsg := "See https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to " +
		"retrieve a valid ID token."

	if t.IssuedAt > now {
		return xerrors.Errorf("ID token issued at future timestamp: %d", t.IssuedAt)
	}

	if t.Expires < now {
		return xerrors.Errorf("ID token has expired. Expired at: %d", t.Expires)
	}

	if t.Subject == "" {
		return xerrors.Errorf("ID token has empty 'sub' (subject) claim. %s", verifyTokenMsg)
	}

	if len(t.Subject) > 128 {
		return xerrors.Errorf("ID token has a 'sub' (subject) claim longer than 128 characters. %s", verifyTokenMsg)
	}

	return nil
}
