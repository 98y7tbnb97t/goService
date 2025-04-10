package handlers

import (
	"echoServer/internal/userService/service"
	"echoServer/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	service service.UserService
}

func RegisterUserHandlers(e *echo.Echo, service service.UserService) {
	h := &UserHandlers{service: service}

	e.GET("/users", h.GetUsers)
	e.POST("/users", h.CreateUser)
	e.PATCH("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}

func (h *UserHandlers) GetUsers(c echo.Context) error {
	page := 1   // Example: default page number
	limit := 10 // Example: default limit
	users, _, err := h.service.GetUsers(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandlers) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	updatedUser, err := h.service.UpdateUser(uint(id), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandlers) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
