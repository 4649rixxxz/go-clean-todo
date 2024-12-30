package todo

import (
	todoDomain "go-clean-todo/domain/todo"
	"go-clean-todo/usecase"
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

func (uc *CreateTodoUsecase) Run(inputDTO CreateTodoUsecaseInputDTO) (*CreateTodoUsecaseOutputDTO, usecase.UsecaseErrorI) {
	todo, todoErr := todoDomain.NewTodo(
		inputDTO.UserID,
		inputDTO.Title,
		inputDTO.Description,
	)
	if todoErr != nil {
		return nil, usecase.NewInvalidInputError(todoErr.Field(), todoErr.Error())
	}
	if createErr := uc.todoRepo.CreateTodo(todo); createErr != nil {
		return nil, usecase.NewInternalServerError("todoの新規作成に失敗しました。")
	}
	outputDTO := &CreateTodoUsecaseOutputDTO{
		TodoID:           todo.TodoID(),
		UserID:           todo.UserID(),
		Title:            todo.Title(),
		Description:      todo.Description(),
		AttachedFilePath: todo.AttachedFilePath(),
		CompletedAt:      todo.CompletedAt(),
		CreatedAt:        todo.CreatedAt(),
		UpdatedAt:        todo.UpdatedAt(),
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
