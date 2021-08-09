package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	wrapper "github.com/rs/cors/wrapper/gin"
)

// NewGinMiddleware - gin-gonic/gin 用のCORSミドルウェアを生成
func NewGinMiddleware() (gin.HandlerFunc, error) {
	options := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Forwarded-For",
			"X-Forwarded-Proto",
			"X-Real-Ip",
		},
		AllowCredentials:   false,
		MaxAge:             1440, // 60m * 24h
		OptionsPassthrough: true,
		Debug:              false,
	}

	return wrapper.New(options), nil
}
