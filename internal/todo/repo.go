package todo

type TodoRepository interface {
	CreateTodo(todo *Todo) error
	FindAll() ([]*Todo, error)
}
