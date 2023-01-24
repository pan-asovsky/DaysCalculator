package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/pan-asovsky/DaysCalculator/internal/services"
	"net/http"
)

func WhenYearRouteHandler(context *gin.Context) {

	year := context.Param("year")
	receivedDate, err := services.GetDateFromString(year)
	if err != nil {
		context.String(http.StatusBadRequest, "Invalid date")
		return
	}

	difference := services.GetDateDifference(receivedDate)
	context.String(http.StatusOK, difference)
}
