package util

import (
	"context"

	"github.com/calmato/gran-book/api/pkg/metadata"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gmd "google.golang.org/grpc/metadata"
)

func SetMetadata(ctx *gin.Context) context.Context {
	md := metadata.GrpcMetadata{
		Ctx: metadata.GinContextToContext(ctx),
	}

	token := ctx.GetHeader("Authorization")
	if token != "" {
		md.Ctx = gmd.AppendToOutgoingContext(md.Ctx, "Authorization", token)
	}

	userID := ctx.GetHeader("userId")
	if userID != "" {
		md.Ctx = gmd.AppendToOutgoingContext(md.Ctx, "userId", userID)
	}

	role := ctx.GetHeader("role")
	if role != "" {
		md.Ctx = gmd.AppendToOutgoingContext(md.Ctx, "role", role)
	}

	requestID := ctx.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}
	md.Ctx = gmd.AppendToOutgoingContext(md.Ctx, "X-Request-ID", requestID)

	return md.Ctx
}
