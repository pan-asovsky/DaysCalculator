package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func WhenYearRouteHandler(context *gin.Context) {

	year := context.Param("year")
	yearAsInt, err := strconv.Atoi(year)

	if err != nil {
		log.Println("Error converting string to int:", err)
	}

	dateString := "01/01/" + year
	dateLayout := "01/01/2006"
	parsedDate, err := time.Parse(dateLayout, dateString)

	if err != nil {
		log.Println("Error parse date:", err)
		context.String(http.StatusBadRequest, "Invalid date")
		return
	}

	currentTime := time.Now()
	currentYear := currentTime.Year()

	daysUntil := parsedDate.Sub(currentTime).Hours() / 24
	daysPassed := currentTime.Sub(parsedDate).Hours() / 24

	if currentYear == yearAsInt && currentTime.Month() == parsedDate.Month() && currentTime.Day() == parsedDate.Day() {
		context.String(http.StatusOK, "Today 1st january")
	} else if currentYear < yearAsInt {
		context.String(http.StatusOK, "Days left: "+float64ToString(daysUntil))
	} else {
		context.String(http.StatusOK, "Days gone: "+float64ToString(math.Abs(daysPassed)))
	}

}

func float64ToString(num float64) string {
	return strconv.Itoa(int(num))
}
