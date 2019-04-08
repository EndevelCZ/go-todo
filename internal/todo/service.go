package todo

import (
	"github.com/sirupsen/logrus"
)

type TodoService interface {
	CreateTodo(todo *Todo) error
	FindAllTodos() ([]*Todo, error)
}
type todoService struct {
	repo TodoRepository
}

// NewTodoService create new TodoService
func NewTodoService(repo TodoRepository) TodoService {
	return &todoService{
		repo,
	}
}
func (s *todoService) FindAllTodos() ([]*Todo, error) {
	todos, err := s.repo.FindAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all todos")
		return nil, err
	}
	logrus.Info("Found all tickets")
	return todos, nil
}
func (s *todoService) CreateTodo(todo *Todo) error {
	// todo.ID = uuid.New().String()
	// todo.Done = false
	// todo.ID = uuid.New()
	todo.Done = false

	if err := s.repo.CreateTodo(todo); err != nil {
		logrus.WithField("error", err).Error("Error creating todo")
		return err
	}
	logrus.WithField("id", todo.ID).Info("Created new todo")
	return nil
}
