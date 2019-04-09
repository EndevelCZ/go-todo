package todo

import (
	"github.com/sirupsen/logrus"
)

type TodoService interface {
	CreateTodo(todo *Todo) error
	FindAllTodos() ([]*Todo, error)
	CheckTodo(int64) (*Todo, error)
	DeleteTodo(int64) (*Todo, error)
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
	todo.Done = false
	if err := s.repo.CreateTodo(todo); err != nil {
		logrus.WithField("error", err).Error("Error creating todo")
		return err
	}
	logrus.WithField("id", todo.ID).Info("Created new todo")
	return nil
}
func (s *todoService) CheckTodo(id int64) (*Todo, error) {
	todo, err := s.repo.CheckTodo(id)
	if err != nil {
		logrus.WithField("error", err).Error("Error checking todo")
		return nil, err
	}
	return todo, nil
}
func (s *todoService) DeleteTodo(id int64) (*Todo, error) {
	todo, err := s.repo.DeleteTodo(id)
	if err != nil {
		logrus.WithField("error", err).Error("Error deleting todo")
		return nil, err
	}
	return todo, nil
}
