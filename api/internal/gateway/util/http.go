package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx.Set("X-Request-ID", requestID)
		ctx.Next()
	}
}
