package todo

type CreateTodoParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
