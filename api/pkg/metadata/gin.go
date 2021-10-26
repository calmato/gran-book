package metadata

import (
	"context"

	"github.com/gin-gonic/gin"
)

type key int

// GinContext - Context変換用のキー
var GinContext key

// GinContextToContext - gin.Context -> context.Context
func GinContextToContext(ctx *gin.Context) context.Context {
	return context.WithValue(ctx.Request.Context(), GinContext, ctx)
}

// GinContextFromContext - context.Context -> gin.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContext)
	if ginContext == nil {
		return nil, ErrNotRetrieveContext
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, ErrInvalidContext
	}

	return gc, nil
}
