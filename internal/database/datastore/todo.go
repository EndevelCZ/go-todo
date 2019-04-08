package datastore

import (
	"context"
	"fmt"

	"github.com/adamplansky/todo/internal/todo"

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
func (db *todoRepository) CreateTodo(todo *todo.Todo) (err error) {
	k := datastore.IncompleteKey("Todo", nil)
	ctx := context.Background()
	k, err = db.connection.Put(ctx, k, todo)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put Book: %v", err)
	}
	return nil
}

func (db *todoRepository) FindAll() ([]*todo.Todo, error) {
	ctx := context.Background()
	todos := make([]*todo.Todo, 0)
	q := datastore.NewQuery("Todo")

	keys, err := db.connection.GetAll(ctx, q, &todos)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list todos: %v", err)
	}

	for i, k := range keys {
		todos[i].ID = k.ID
	}
	return todos, nil
}
