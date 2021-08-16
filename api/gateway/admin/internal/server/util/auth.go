package util

import (
	"fmt"
	"strings"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	"github.com/calmato/gran-book/api/gateway/admin/pkg/firebase/authentication"
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Authenticator interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
	HasAdminRole() gin.HandlerFunc
}

type authenticator struct {
	auth *authentication.Auth
	api  pb.AuthServiceClient
}

func NewAuthenticator(fa *authentication.Auth, authConn *grpc.ClientConn) Authenticator {
	ac := pb.NewAuthServiceClient(authConn)

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

		ctx.Next()
	}
}

func (a *authenticator) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		out, err := a.api.GetAuth(ctx, &pb.Empty{})
		if err != nil {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		role := entity.Role(out.GetRole())
		if role == entity.RoleUser {
			ErrorHandling(ctx, entity.ErrForbidden.New(err))
			return
		}

		ctx.Next()
	}
}

func (a *authenticator) HasAdminRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		out, err := a.api.GetAuth(ctx, &pb.Empty{})
		if err != nil {
			ErrorHandling(ctx, entity.ErrUnauthenticated.New(err))
			return
		}

		role := entity.Role(out.GetRole())
		if role != entity.RoleAdmin {
			ErrorHandling(ctx, entity.ErrForbidden.New(err))
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
