package taskService

import (
	"errors"
)

var (
	ErrTaskNotFound = errors.New("Task not found")
)

type TaskService struct {
	repo TaskRepository
}

func NewService(repo *taskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {

	return s.repo.DeleteTaskByID(id)
}
