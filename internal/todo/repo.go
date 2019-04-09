package todo

type TodoRepository interface {
	CreateTodo(todo *Todo) error
	FindAll() ([]*Todo, error)
	CheckTodo(int64) (*Todo, error)
	DeleteTodo(int64) (*Todo, error)
}
