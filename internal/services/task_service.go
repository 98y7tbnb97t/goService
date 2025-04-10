package services

import (
	"echoServer/internal/repositories"
	"echoServer/models"
)

type Task = models.Task

// TaskService defines methods for task-related operations.
type TaskService interface {
	GetTasks() ([]Task, error)
	CreateTask(task *Task) error
	UpdateTask(id string, task *Task) error
	PatchTask(id string, task *Task) error
	DeleteTask(id string) error
	GetTaskByID(id string, task *Task) error
}

// taskServiceImpl is the concrete implementation of TaskService.
type taskServiceImpl struct {
	repo repositories.TaskRepository
}

// NewService creates a new instance of TaskService.
func NewService(repo repositories.TaskRepository) TaskService {
	return &taskServiceImpl{repo: repo}
}

func (s *taskServiceImpl) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *taskServiceImpl) CreateTask(task *Task) error {
	return s.repo.CreateTask(task)
}

func (s *taskServiceImpl) UpdateTask(id string, task *Task) error {
	return s.repo.UpdateTask(id, task)
}

func (s *taskServiceImpl) PatchTask(id string, task *Task) error {
	return s.repo.PatchTask(id, task)
}

func (s *taskServiceImpl) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func (s *taskServiceImpl) GetTaskByID(id string, task *Task) error {
	return s.repo.GetTaskByID(id, task)
}
