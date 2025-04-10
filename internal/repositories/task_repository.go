package repositories

import (
	"echoServer/db"
	"echoServer/models"

	"gorm.io/gorm"
)

type Task = models.Task

// TaskRepository defines the interface for task repository operations.
type TaskRepository interface {
	GetTasks() ([]Task, error)
	CreateTask(task *Task) error
	UpdateTask(id string, task *Task) error
	PatchTask(id string, task *Task) error
	DeleteTask(id string) error
	GetTaskByID(id string, task *Task) error
}

// GormTaskRepository is the GORM-based implementation of TaskRepository.
type GormTaskRepository struct {
	db *gorm.DB
}

// NewGormRepository creates a new instance of GormTaskRepository.
func NewGormRepository(db *gorm.DB) TaskRepository {
	return &GormTaskRepository{db: db}
}

// Default repository instance using the global DB
var DefaultTaskRepository TaskRepository = NewGormRepository(db.DB)

// GetTasks retrieves all tasks from the database.
func GetTasks() ([]Task, error) {
	return DefaultTaskRepository.GetTasks()
}

// CreateTask creates a new task in the database.
func CreateTask(task *Task) error {
	return DefaultTaskRepository.CreateTask(task)
}

// UpdateTask updates an existing task in the database.
func UpdateTask(id string, task *Task) error {
	return DefaultTaskRepository.UpdateTask(id, task)
}

// PatchTask partially updates an existing task in the database.
func PatchTask(id string, task *Task) error {
	return DefaultTaskRepository.PatchTask(id, task)
}

// DeleteTask removes a task from the database.
func DeleteTask(id string) error {
	return DefaultTaskRepository.DeleteTask(id)
}

// GetTaskByID retrieves a task by its ID from the database.
func GetTaskByID(id string, task *Task) error {
	return DefaultTaskRepository.GetTaskByID(id, task)
}

func (r *GormTaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *GormTaskRepository) CreateTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *GormTaskRepository) UpdateTask(id string, task *Task) error {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}
	return r.db.Save(task).Error
}

func (r *GormTaskRepository) PatchTask(id string, task *Task) error {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}
	return r.db.Save(task).Error
}

func (r *GormTaskRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, id).Error
}

func (r *GormTaskRepository) GetTaskByID(id string, task *Task) error {
	return r.db.First(task, "id = ? AND deleted_at IS NULL", id).Error
}
