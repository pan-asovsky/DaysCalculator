package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	in "github.com/pan-asovsky/DaysCalculator/internal/transport"
	"net/http"
	"net/http/httptest"
	conv "strconv"
	"testing"
	"time"
)

func TestWhenYearRouteHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/when/:year", in.WhenYearRouteHandler)

	currentTime := time.Now()
	currentYear := currentTime.Year()
	regex := "\\w+ \\w+[:] \\d+"

	t.Run("CurrentYear", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear), nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.MatchRegex(t, resp.Body.String(), regex)
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("FutureYear", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear+1), nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.MatchRegex(t, resp.Body.String(), regex)
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("PastYear", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/when/"+conv.Itoa(currentYear-1), nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.MatchRegex(t, resp.Body.String(), regex)
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, received %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("InvalidDate", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/when/abc", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, "Invalid date", resp.Body.String())
		if resp.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, received %d", http.StatusBadRequest, resp.Code)
		}
	})

}
