package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func main() {

	router := gin.Default()
	router.Use(XPingHeaderMiddleware())
	router.GET("/when/:year", whenYearRouteHandler)

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run(): ", err)
	}
}

func XPingHeaderMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if ping := context.GetHeader("X-PING"); ping == "ping" {
			context.Header("X-PONG", "pong")
		}
		context.Next()
	}
}

func whenYearRouteHandler(context *gin.Context) {

	year := context.Param("year")
	yearAsInt, err := strconv.Atoi(year)

	if err != nil {
		log.Println("Error converting string to int: ", err)
	}

	dateString := "01/01/" + year
	dateLayout := "01/01/2006"
	parsedDate, err := time.Parse(dateLayout, dateString)

	if err != nil {
		log.Println("Error parse date: ", err)
		context.String(http.StatusBadRequest, "invalid date")
		return
	}

	currentTime := time.Now()
	currentYear := currentTime.Year()

	daysUntil := parsedDate.Sub(currentTime).Hours() / 24
	daysPassed := currentTime.Sub(parsedDate).Hours() / 24

	if currentYear == yearAsInt && currentTime.Month() == parsedDate.Month() && currentTime.Day() == parsedDate.Day() {
		context.String(http.StatusOK, "Today 1st january")
	} else if currentYear < yearAsInt {
		context.String(http.StatusOK, "Days left: "+IntToString(daysUntil))
	} else {
		context.String(http.StatusOK, "Days gone: "+IntToString(math.Abs(daysPassed)))
	}

}

func IntToString(num float64) string {
	return strconv.Itoa(int(num))
}
