package http_gin_main

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	//assert.Equal(t, "pong", w.Body.String())
}
