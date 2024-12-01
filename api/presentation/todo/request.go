package todo

type CreateTodoParams struct {
	Title       string `json:"title" validate:"required" validate:"title is required"`
	Description string `json:"description" validate:"required"`
}
