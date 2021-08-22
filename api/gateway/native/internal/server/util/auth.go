package util

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"github.com/gin-gonic/gin"
)

type Authenticator interface {
	Authentication() gin.HandlerFunc
}

type authenticator struct {
	auth *authentication.Auth
}

func NewAuthenticator(fa *authentication.Auth) Authenticator {
	return &authenticator{
		auth: fa,
	}
}

func (a *authenticator) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := a.getToken(ctx)
		if err != nil {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		userID, err := a.auth.VerifyIDToken(ctx, token)
		if err != nil || userID == "" {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		a.setToken(ctx, token)
		a.setAuth(ctx, userID)

		ctx.Next()
	}
}

func (a *authenticator) setAuth(ctx *gin.Context, userID string) {
	ctx.Set("userId", userID)
}

func (a *authenticator) setToken(ctx *gin.Context, token string) {
	ctx.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (a *authenticator) getToken(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", fmt.Errorf("authorization header is not contain")
	}

	t := strings.Replace(token, "Bearer ", "", 1)
	return t, nil
}
