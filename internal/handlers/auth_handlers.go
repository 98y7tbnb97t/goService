package handlers

import (
	"net/http"

	"echoServer/config"
	"echoServer/internal/auth"
	"echoServer/internal/middleware"
	"echoServer/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	cfg         *config.Config
	userService interface {
		Authenticate(email, password string) (models.User, error)
	}
	tokenService *auth.TokenService
}

func RegisterAuthHandlers(e *echo.Echo, cfg *config.Config, us interface {
	Authenticate(email, password string) (models.User, error)
}, ts *auth.TokenService) {
	h := &AuthHandler{cfg: cfg, userService: us, tokenService: ts}

	e.POST("/login", h.Login)
	e.POST("/refresh", h.Refresh)
	e.POST("/logout", h.Logout, middleware.AuthMiddleware)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.userService.Authenticate(creds.Email, creds.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	accessToken, err := h.tokenService.GenerateAccessToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate access token"})
	}

	refreshToken, err := h.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate refresh token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    h.cfg.JWT.AccessDuration * 60,
	})
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	newToken, err := h.tokenService.Refresh(c.Request())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": newToken})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	if err := h.tokenService.RevokeToken(claims["jti"].(string)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to logout"})
	}

	return c.NoContent(http.StatusNoContent)
}
