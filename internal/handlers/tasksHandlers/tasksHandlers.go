package tasksHandlers

import (
	"echoServer/internal/services"
	"net/http"
	"time"

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
	if task.UserID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "field user_id is required"})
	}
	// Создаём задачу с привязкой к пользователю
	if err := services.CreateTaskForUser(task.UserID, &task); err != nil {
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
	var task services.Task
	if err := services.GetTaskByID(id, &task); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	task.DeletedAt = time.Now()
	if err := services.UpdateTask(id, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "task deleted"})
}
