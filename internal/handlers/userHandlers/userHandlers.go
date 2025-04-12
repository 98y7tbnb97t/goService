package userHandlers

import (
	"net/http"
	"strconv"

	"echoServer/internal/services"

	"github.com/labstack/echo/v4"
)

// RegisterUserHandlers registers user endpoints on the Echo router.
func RegisterUserHandlers(e *echo.Echo) {
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.DELETE("/users/:id", deleteUser)
	e.PATCH("/users/:id", patchUser)
}

// getUsers handles GET /users to fetch all users.
func getUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

// createUser handles POST /users to create a new user.
func createUser(c echo.Context) error {
	var user services.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := services.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

// deleteUser handles DELETE /users/:id to remove an existing user.
func deleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	// Check if user exists before deletion
	user, err := services.GetUserByID(id)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	if err := services.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "user deleted"})
}

// patchUser handles PATCH /users/:id to update an existing user partially.
func patchUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	var updates services.User
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Ensure the user exists before patching
	user, err := services.GetUserByID(id)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	if err := services.UpdateUser(id, &updates); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
	}

	updatedUser, err := services.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch updated user"})
	}
	return c.JSON(http.StatusOK, updatedUser)
}
