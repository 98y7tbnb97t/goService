package services

import (
	"echoServer/internal/repositories"
	"echoServer/models"
)

type Task = models.Task

func GetTasks() ([]Task, error) {
	return repositories.GetTasks()
}

func CreateTask(task *Task) error {
	return repositories.CreateTask(task)
}

func CreateTaskForUser(userID uint, task *Task) error {
	task.UserID = userID
	return repositories.CreateTask(task)
}

func UpdateTask(id string, task *Task) error {
	return repositories.UpdateTask(id, task)
}

func PatchTask(id string, task *Task) error {
	return repositories.PatchTask(id, task)
}

func DeleteTask(id string) error {
	return repositories.DeleteTask(id)
}

func GetTaskByID(id string, task *Task) error {
	return repositories.GetTaskByID(id, task)
}
