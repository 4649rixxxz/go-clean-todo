package todo

type TodoRepository interface {
	CreateTodo(todo *Todo) (*Todo, error)
}
