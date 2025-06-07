package repositories

import (
	"echoServer/db"
	"echoServer/models"
)

type Task = models.Task

func GetTasks() ([]Task, error) {
	var tasks []Task
	result := db.DB.Omit("created_at", "updated_at", "deleted_at").Find(&tasks)
	return tasks, result.Error
}

func CreateTask(task *Task) error {
	return db.DB.Create(task).Error
}

func UpdateTask(id string, task *Task) error {
	var existingTask Task
	if err := db.DB.First(&existingTask, id).Error; err != nil {
		return err
	}
	return db.DB.Save(task).Error
}

func PatchTask(id string, task *Task) error {
	var existingTask Task
	if err := db.DB.First(&existingTask, id).Error; err != nil {
		return err
	}
	return db.DB.Save(task).Error
}

func DeleteTask(id string) error {
	return db.DB.Delete(&Task{}, id).Error
}

func GetTaskByID(id string, task *Task) error {
	// Omit the _at fields so they are not returned
	return db.DB.Omit("created_at", "updated_at", "deleted_at").First(task, "id = ? AND deleted_at IS NULL", id).Error
}

// GetTasksForUser retrieves all tasks belonging to a specific user
func GetTasksForUser(userID uint) ([]Task, error) {
	var tasks []Task
	// Omit the _at fields from the result
	result := db.DB.Omit("created_at", "updated_at", "deleted_at").Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}
