package server

import "github.com/gin-gonic/gin"

func XPingHeaderMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if ping := context.GetHeader("X-PING"); ping == "ping" {
			context.Header("X-PONG", "pong")
		}
		context.Next()
	}
}
