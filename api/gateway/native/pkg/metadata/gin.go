package metadata

import (
	"context"
	"fmt"

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
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}
