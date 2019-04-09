package todo

import (
	"context"

	"github.com/adamplansky/todo/pb"
	"github.com/sirupsen/logrus"
)

type TodoGrpcHandler interface {
	AddTodo(ctx context.Context, text *pb.Text) (*pb.Todo, error)
	ListTodos(ctx context.Context, void *pb.Void) (*pb.TodoList, error)
	CheckTodo(context.Context, *pb.Integer) (*pb.Todo, error)
	DeleteTodo(context.Context, *pb.Integer) (*pb.Todo, error)
}

type todoGrpcHandler struct {
	service TodoService
}

func NewTodoGrpcHandler(service TodoService) TodoGrpcHandler {
	return &todoGrpcHandler{
		service,
	}
}

func (s *todoGrpcHandler) AddTodo(ctx context.Context, text *pb.Text) (*pb.Todo, error) {
	todo := &Todo{
		Text: text.Text,
	}
	err := s.service.CreateTodo(todo)
	if err != nil {
		logrus.WithField("error", err).Error("Error creating todo")
		return nil, err
	}
	return todoToPbTodo(todo), nil
	// x := &todoPbSerializer{handler: s}
	// return x.Serialize(todo), nil
}

func (s *todoGrpcHandler) ListTodos(ctx context.Context, void *pb.Void) (*pb.TodoList, error) {
	todos, err := s.service.FindAllTodos()
	if err != nil {
		logrus.WithField("error", err).Error("Error finding all todos")
		return nil, err
	}
	// type TodoList struct {
	// 	Todos
	list := &pb.TodoList{}
	for _, t := range todos {
		list.Todos = append(list.Todos, todoToPbTodo(t))
	}

	return list, nil
}

func (s *todoGrpcHandler) CheckTodo(ctx context.Context, id *pb.Integer) (*pb.Todo, error) {
	todo, err := s.service.CheckTodo(id.Id)
	if err != nil {
		logrus.WithField("error", err).Error("Error update todo")
		return nil, err
	}
	return todoToPbTodo(todo), nil

}

func (s *todoGrpcHandler) DeleteTodo(ctx context.Context, id *pb.Integer) (*pb.Todo, error) {
	todo, err := s.service.DeleteTodo(id.Id)
	if err != nil {
		logrus.WithField("error", err).Error("Error update todo")
		return nil, err
	}
	return todoToPbTodo(todo), nil
}
