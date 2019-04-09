package datastore

import (
	"context"
	"fmt"

	"github.com/adamplansky/todo/internal/todo"
	"github.com/sirupsen/logrus"

	"cloud.google.com/go/datastore"
)

type todoRepository struct {
	connection *datastore.Client
}

// Ensure datastoreDB conforms to the BookDatabase interface.
// var _ todo.TodoRepository = &todoRepository{}

// NewDataStoreTodoRepository return todoRepository
func NewDataStoreTodoRepository(connection *datastore.Client) todo.TodoRepository {
	return &todoRepository{
		connection,
	}
}
func (db *todoRepository) datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey("Todo", id, nil)
}
func (db *todoRepository) CreateTodo(todo *todo.Todo) (err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("Todo", nil)
	k, err = db.connection.Put(ctx, k, todo)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put Todo: %v", err)
	}
	todo.ID = k.ID
	return nil
}

func (db *todoRepository) FindAll() ([]*todo.Todo, error) {
	logrus.Info("Proccessing FindAll datastore")
	ctx := context.Background()
	todos := make([]*todo.Todo, 0)
	q := datastore.NewQuery("Todo")
	logrus.Info("Proccessing get all")
	keys, err := db.connection.GetAll(ctx, q, &todos)
	logrus.Info("Processed get all")
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list todos: %v", err)
	}
	for i, k := range keys {
		todos[i].ID = k.ID
	}
	logrus.Info("returning todos")
	return todos, nil
}
func (db *todoRepository) CheckTodo(id int64) (*todo.Todo, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	var todo todo.Todo
	if err := db.connection.Get(ctx, k, &todo); err != nil {
		return nil, fmt.Errorf("datastoredb: could not delete [get] todo: %v", err)
	}
	todo.Done = !todo.Done
	if _, err := db.connection.Put(ctx, k, &todo); err != nil {
		return nil, fmt.Errorf("datastoredb: could not update todo: %v", err)
	}
	logrus.Info("key has been updated", todo)
	return &todo, nil
}
func (db *todoRepository) DeleteTodo(id int64) (*todo.Todo, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	var todo todo.Todo
	if err := db.connection.Get(ctx, k, &todo); err != nil {
		return nil, fmt.Errorf("datastoredb: could not delete [get]todo: %v", err)
	}
	err := db.connection.Delete(ctx, k)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not delete todo: %v", err)
	}
	logrus.Info("key has been deleted", todo)
	return &todo, nil
}
