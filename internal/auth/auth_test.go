package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"echoServer/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// mockTokenService implements TokenServiceInterface for testing.
type mockTokenService struct {
	// you can add fields to simulate different behavior if needed
}

func (m *mockTokenService) Refresh(r *http.Request) (string, error) {
	// For test purposes, if the header is set to "bad", return error.
	if r.Header.Get("X-Refresh-Token") == "bad" {
		return "", errors.New("invalid refresh token")
	}
	return "test-new-token", nil
}

func TestRefreshHandlerSuccess(t *testing.T) {
	e := echo.New()
	// Create a dummy request with X-Refresh-Token header.
	req := httptest.NewRequest(http.MethodGet, "/refresh", nil)
	req.Header.Set("X-Refresh-Token", "good")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create handler instance with the mock token service.
	h := &handlers.AuthHandler{tokenService: &mockTokenService{}}

	// Invoke the Refresh handler.
	err := h.Refresh(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Check response token.
	var response map[string]string
	err = json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "test-new-token", response["token"])
}

func TestRefreshHandlerFailure(t *testing.T) {
	e := echo.New()
	// Create a dummy request with a missing or bad token.
	req := httptest.NewRequest(http.MethodGet, "/refresh", nil)
	req.Header.Set("X-Refresh-Token", "bad")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &handlers.AuthHandler{tokenService: &mockTokenService{}}

	err := h.Refresh(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	var response map[string]string
	err = json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "invalid refresh token")
}
