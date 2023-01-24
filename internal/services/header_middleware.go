package services

import "github.com/gin-gonic/gin"

const requestHeader = "X-PING"
const requestHeaderValue = "ping"

const responseHeader = "X-PONG"
const responseHeaderValue = "pong"

func XPingHeaderMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if ping := context.GetHeader(requestHeader); ping == requestHeaderValue {
			context.Header(responseHeader, responseHeaderValue)
		}
		context.Next()
	}
}
