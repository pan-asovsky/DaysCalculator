package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/pan-asovsky/DaysCalculator/internal/services"
)

const relativePath = "/when/:year"

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(services.XPingHeaderMiddleware())
	router.GET(relativePath, WhenYearRouteHandler)

	return router
}
