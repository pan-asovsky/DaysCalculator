package internal

import "github.com/gin-gonic/gin"

const relativePath = "/when/:year"

func GetRouter() *gin.Engine {

	router := gin.Default()
	router.Use(XPingHeaderMiddleware())
	router.GET(relativePath, WhenYearRouteHandler)

	return router
}
