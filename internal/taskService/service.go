package taskService

type TaskService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) DeleteTask(task Task) error {
	return s.repo.DeleteTaskByID(task.ID)
}

//func (s *taskService) UpdateTask(id uint) (Task, error) {
//	return s.repo.UpdateTaskByID()
