package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler interface {
	HealthCheck(ctx *gin.Context)
	SetRequestID() gin.HandlerFunc
}

type httpHandler struct{}

func NewHTTPHandler() HTTPHandler {
	return &httpHandler{}
}

func (h *httpHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *httpHandler) SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewUUID()
		}

		ctx.Set("X-Request-ID", requestID)
		ctx.Next()
	}
}
