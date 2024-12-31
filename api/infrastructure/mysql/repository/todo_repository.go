package repository

import (
	todoDomain "go-clean-todo/domain/todo"
	"go-clean-todo/infrastructure/mysql"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository() todoDomain.TodoRepository {
	return &todoRepository{
		db: mysql.GetDB(),
	}
}

func (r *todoRepository) CreateTodo(todo *todoDomain.Todo) error {
	todoORM := mysql.Todo{
		UserID:      todo.UserID(),
		Title:       todo.Title(),
		Description: todo.Description(),
	}
	if err := r.db.Create(&todoORM).Error; err != nil {
		return err
	}

	todo.Set(
		todoORM.TodoID,
		todoORM.UserID,
		todoORM.Title,
		todoORM.Description,
		todoORM.AttachedFilePath,
		todoORM.CompletedAt,
		todoORM.CreatedAt,
		todoORM.UpdatedAt,
	)

	return nil
}
