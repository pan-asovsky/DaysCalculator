package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pan-asovsky/DaysCalculator/internal"
	"log"
)

func main() {

	router := gin.Default()
	router.Use(internal.XPingHeaderMiddleware())
	router.GET("/when/:year", internal.WhenYearRouteHandler)

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run():", err)
	}
}
