package handlers

import (
	"echoServer/internal/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	ts services.TaskService
}

// RegisterTaskHandlers registers task routes, accepts a TaskService instance
func RegisterTaskHandlers(e *echo.Echo, ts services.TaskService) {
	h := &TaskHandler{ts: ts}
	e.GET("/tasks", h.getTasks)
	e.POST("/tasks", h.createTask)
	e.PUT("/tasks/:id", h.updateTask)
	e.PATCH("/tasks/:id", h.patchTask)
	e.DELETE("/tasks/:id", h.deleteTask)
}

func (h *TaskHandler) getTasks(c echo.Context) error {
	tasks, err := h.ts.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) createTask(c echo.Context) error {
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.ts.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) updateTask(c echo.Context) error {
	id := c.Param("id")
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.ts.UpdateTask(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) patchTask(c echo.Context) error {
	id := c.Param("id")
	var task services.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.ts.PatchTask(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) deleteTask(c echo.Context) error {
	id := c.Param("id")
	var task services.Task
	if err := h.ts.GetTaskByID(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	task.DeletedAt = time.Now()
	if err := h.ts.UpdateTask(id, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "task deleted"})
}
