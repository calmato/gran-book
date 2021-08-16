package util

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
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

		ctx.Next()
	}
}

func (a *authenticator) getToken(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", fmt.Errorf("authorization header is not contain")
	}

	t := strings.Replace(token, "Bearer ", "", 1)
	return t, nil
}
