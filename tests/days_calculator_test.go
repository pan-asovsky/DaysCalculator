package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/pan-asovsky/DaysCalculator/pkg/server"
	"net/http"
	"net/http/httptest"
	conv "strconv"
	"testing"
	"time"
)

func TestWhenYearRouteHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", server.WhenYearRouteHandler)

	currentTime := time.Now()
	currentYear := currentTime.Year()

	// Current
	req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	regex := "\\w+ \\w+[:] \\d+"
	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}

	// Future
	req, _ = http.NewRequest("GET", "/when/"+conv.Itoa(currentYear+1), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}

	// Past
	req, _ = http.NewRequest("GET", "/when/"+conv.Itoa(currentYear-1), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.MatchRegex(t, resp.Body.String(), regex)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
	}

	// Invalid
	req, _ = http.NewRequest("GET", "/when/abc", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "Invalid date", resp.Body.String())
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, received %d", http.StatusBadRequest, resp.Code)
	}

}
