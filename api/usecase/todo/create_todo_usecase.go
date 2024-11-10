package todo

import (
	todoDomain "go-clean-todo/domain/todo"
	"time"
)

type CreateTodoUsecase struct {
	todoRepo todoDomain.TodoRepository
}

func NewCreateTodoUsecase(
	todoRepo todoDomain.TodoRepository,
) *CreateTodoUsecase {
	return &CreateTodoUsecase{
		todoRepo: todoRepo,
	}
}

func (uc *CreateTodoUsecase) Run(inputDTO CreateTodoUsecaseInputDTO) (*CreateTodoUsecaseOutputDTO, error) {
	todo, todoErr := todoDomain.NewTodo(
		inputDTO.UserID,
		inputDTO.Title,
		inputDTO.Description,
	)
	if todoErr != nil {
		return nil, todoErr
	}
	createdTodo, createErr := uc.todoRepo.CreateTodo(todo)
	if createErr != nil {
		return nil, createErr
	}
	outputDTO := &CreateTodoUsecaseOutputDTO{
		TodoID:           createdTodo.TodoID(),
		UserID:           createdTodo.UserID(),
		Title:            createdTodo.Title(),
		Description:      createdTodo.Description(),
		AttachedFilePath: createdTodo.AttachedFilePath(),
		CompletedAt:      createdTodo.CompletedAt(),
		CreatedAt:        createdTodo.CreatedAt(),
		UpdatedAt:        createdTodo.UpdatedAt(),
	}

	return outputDTO, nil
}

type CreateTodoUsecaseInputDTO struct {
	UserID      uint
	Title       string
	Description string
}

type CreateTodoUsecaseOutputDTO struct {
	TodoID           uint
	UserID           uint
	Title            string
	Description      string
	AttachedFilePath *string
	CompletedAt      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
