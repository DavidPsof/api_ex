package apiserver

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandlePing(t *testing.T) {
	s := NewAPIServer(NewConfig())
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	s.HandlePing().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "pong")
}
