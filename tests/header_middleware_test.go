package tests

import (
	"github.com/gin-gonic/gin"
	in "github.com/pan-asovsky/DaysCalculator/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXPingHeaderMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(in.XPingHeaderMiddleware())

	t.Run("HeaderIsSet", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Add("X-PING", "ping")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Header().Get("X-PONG") != "pong" {
			t.Error("The expected X-PONG header should be set to 'pong'")
		}
	})

	t.Run("HeaderIsUnset", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Header().Get("X-PONG") == "pong" {
			t.Error("The expected X-PONG header should not be set")
		}
	})

}
