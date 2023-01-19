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
	router.GET("/when/:year", whenYearRouteHandler)

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run(): ", err)
	}
}

func whenYearRouteHandler(context *gin.Context) {

	year := context.Param("year")
	yearAsInt, err := strconv.Atoi(year)

	if err != nil {
		log.Println("Error converting string to int: ", err)
	}

	// 2006 используется как формат даты для функции time.Parse().
	// При изменении на другой год, формат перестаёт соответствовать и выбрасывается ошибка.
	dateString := "01/01/" + year
	dateLayout := "01/01/2006"
	parsedDate, err := time.Parse(dateLayout, dateString)

	if err != nil {
		log.Println("Error parse date: ", err)
		context.String(http.StatusBadRequest, "invalid date")
		return
	}

	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentYear := currentTime.Year()

	daysUntil := parsedDate.Sub(currentTime).Hours() / 24
	daysPassed := currentTime.Sub(parsedDate).Hours() / 24

	if currentYear == yearAsInt && currentMonth == parsedDate.Month() && currentTime.Day() == parsedDate.Day() {
		context.String(http.StatusOK, "Today 1st january")
	} else if currentYear < yearAsInt {
		context.String(http.StatusOK, "Days left: "+strconv.Itoa(int(daysUntil)))
	} else {
		context.String(http.StatusOK, "Days gone: "+strconv.Itoa(int(math.Abs(daysPassed))))
	}

}
