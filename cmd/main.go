package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pan-asovsky/DaysCalculator/pkg/server"
	"log"
)

func main() {

	router := gin.Default()
	router.Use(server.XPingHeaderMiddleware())
	router.GET("/when/:year", server.WhenYearRouteHandler)

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run():", err)
	}
}
