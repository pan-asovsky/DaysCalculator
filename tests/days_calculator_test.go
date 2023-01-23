package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/pan-asovsky/DaysCalculator/internal"
	"net/http"
	"net/http/httptest"
	conv "strconv"
	"testing"
	"time"
)

var currentTime = time.Now()
var currentYear = currentTime.Year()
var regex = "\\w+ \\w+[:] \\d+"

func TestWhenYearRouteHandlerCurrent(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", internal.WhenYearRouteHandler)

	req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}
}

func TestWhenYearRouteHandlerFuture(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", internal.WhenYearRouteHandler)

	req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear+1), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}
}

func TestWhenYearRouteHandlerPast(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", internal.WhenYearRouteHandler)

	req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear-1), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}

}

func TestWhenYearRouteHandlerInvalid(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", internal.WhenYearRouteHandler)

	req, _ := http.NewRequest("GET", "/when/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "Invalid date", resp.Body.String())
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, received %d", http.StatusBadRequest, resp.Code)
	}
}
