package handlers

import (
	"echoServer/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterTaskHandlers(e *echo.Echo) {
	e.GET("/tasks", getTasks)
	e.POST("/tasks", createTask)
	e.PUT("/tasks/:id", updateTask)
	e.PATCH("/tasks/:id", patchTask)
	e.DELETE("/tasks/:id", deleteTask)
}

func getTasks(c echo.Context) error {
	tasks, err := services.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func createTask(c echo.Context) error {
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := services.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}
	return c.JSON(http.StatusOK, task)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := services.UpdateTask(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := services.PatchTask(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := services.DeleteTask(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "task deleted"})
}
