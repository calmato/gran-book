package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler interface {
	Check(ctx *gin.Context)
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
