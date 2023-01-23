package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

	daysUntil := int(parsedDate.Sub(currentTime).Hours() / 24)
	daysPassed := int(currentTime.Sub(parsedDate).Hours() / 24)

	if currentYear == yearAsInt && currentTime.Month() == parsedDate.Month() && currentTime.Day() == parsedDate.Day() {
		context.String(http.StatusOK, "Today 1st january")
	} else if currentYear < yearAsInt {
		context.String(http.StatusOK, fmt.Sprintf("Days left: %d", daysUntil))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Days gone: %d", daysPassed))
	}

}
