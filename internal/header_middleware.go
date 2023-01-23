package internal

import "github.com/gin-gonic/gin"

var responseHeader = "X-PONG"
var responseHeaderValue = "pong"

var requestHeader = "X-PING"
var requestHeaderValue = "ping"

func XPingHeaderMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if ping := context.GetHeader(requestHeader); ping == requestHeaderValue {
			context.Header(responseHeader, responseHeaderValue)
		}
		context.Next()
	}
}
