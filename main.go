package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func addTask(task string, isDone bool) error {
	newTask := Task{
		Task:   task,
		IsDone: isDone,
	}
	result := db.Create(&newTask)
	return result.Error
}

func main() {
	initDB()
	err := addTask("new task", false)
	if err != nil {
		fmt.Println("Error adding task:", err)
	}

	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error {
		var tasks []Task
		db.Find(&tasks)
		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/tasks", func(c echo.Context) error {
		var task Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}
		db.Create(&task)
		return c.JSON(http.StatusOK, task)
	})

	e.PUT("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}

		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		db.Save(&task)
		return c.JSON(http.StatusOK, task)
	})

	e.PATCH("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}

		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		db.Save(&task)
		return c.JSON(http.StatusOK, task)
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
		id := c.Param("id")
		if err := db.Delete(&Task{}, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "task deleted"})
	})

	e.Start(":8882")
}
