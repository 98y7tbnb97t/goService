package services

import (
	"echoServer/internal/repositories"
	"echoServer/models"
	"errors"
)

type Task = models.Task

// validateTask ensures that the task has a valid (non-zero) UserID and that a user with such ID exists.
func validateTask(task *Task) error {
	if task.UserID == 0 {
		return errors.New("UserID must not be 0")
	}
	// Convert task.UserID to int as GetUserByID expects an int.
	if _, err := repositories.GetUserByID(int(task.UserID)); err != nil {
		return errors.New("User not found for provided UserID")
	}
	return nil
}

func GetTasks() ([]Task, error) {
	return repositories.GetTasks()
}

func CreateTask(task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return repositories.CreateTask(task)
}

func CreateTaskForUser(userID uint, task *Task) error {
	task.UserID = userID
	if err := validateTask(task); err != nil {
		return err
	}
	return repositories.CreateTask(task)
}

func UpdateTask(id string, task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return repositories.UpdateTask(id, task)
}

func PatchTask(id string, task *Task) error {
	if err := validateTask(task); err != nil {
		return err
	}
	return repositories.PatchTask(id, task)
}

func DeleteTask(id string) error {
	return repositories.DeleteTask(id)
}

func GetTaskByID(id string, task *Task) error {
	return repositories.GetTaskByID(id, task)
}
