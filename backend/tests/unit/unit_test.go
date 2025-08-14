package unit

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/handlers" // ou onde estiver seu PostUser
)

func TestPostUserHandler(t *testing.T) {
	body := `{"name": ""}` // campo inválido

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.PostUser) // substitua com sua função real
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Esperado status 400, veio %d", rr.Code)
	}
}
