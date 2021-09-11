package util

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/gateway/admin/proto/user"
	"github.com/gin-gonic/gin"
)

type Authenticator interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
	HasAdminRole() gin.HandlerFunc
}

type authenticator struct {
	auth *authentication.Auth
	api  user.AuthServiceClient
}

func NewAuthenticator(fa *authentication.Auth, ac user.AuthServiceClient) Authenticator {
	return &authenticator{
		auth: fa,
		api:  ac,
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
		a.setAuth(ctx, userID, entity.RoleUser)

		ctx.Next()
	}
}

func (a *authenticator) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		out, err := a.api.GetAuth(ctx, &user.Empty{})
		if err != nil {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		current := entity.NewAuth(out.Auth)

		if current.Role() == entity.RoleUser {
			ErrorHandling(ctx, entity.ErrForbidden.New(err))
			return
		}

		a.setAuth(ctx, current.Id, current.Role())

		ctx.Next()
	}
}

func (a *authenticator) HasAdminRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		out, err := a.api.GetAuth(ctx, &user.Empty{})
		if err != nil {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		current := entity.NewAuth(out.Auth)

		if current.Role() != entity.RoleAdmin {
			ErrorHandling(ctx, entity.ErrForbidden.New(err))
			return
		}

		a.setAuth(ctx, current.Id, current.Role())

		ctx.Next()
	}
}

func (a *authenticator) setAuth(ctx *gin.Context, userID string, role entity.Role) {
	if userID != "" {
		ctx.Set("userId", userID)
	}

	if role != entity.RoleUser {
		ctx.Set("role", role)
	}
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
