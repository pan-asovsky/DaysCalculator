package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/pan-asovsky/DaysCalculator/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXPingHeaderMiddlewareSet(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(internal.XPingHeaderMiddleware())

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Add("X-PING", "ping")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Header().Get("X-PONG") != "pong" {
		t.Error("The expected X-PONG header should be set to 'pong'")
	}
}

func TestXPingHeaderMiddlewareUnset(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(internal.XPingHeaderMiddleware())

	req, _ := http.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Header().Get("X-PONG") == "pong" {
		t.Error("The expected X-PONG header should not be set")
	}
}
