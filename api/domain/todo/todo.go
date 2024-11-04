package todo

import (
	"time"
	"unicode/utf8"

	errDomain "go-clean-todo/domain/error"
)

type Todo struct {
	todoID           int
	userID           int
	title            string
	description      string
	attachedFilePath *string
	completedAt      *time.Time
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        *time.Time
}

const (
	titleLengthMin       = 0
	titleLengthMax       = 50
	descriptionLengthMin = 0
	descriptionLengthMax = 200
)

func newTodo(
	todoID int,
	userID int,
	title string,
	description string,
	attachedFilePath *string,
	completedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*Todo, error) {
	titleLength := utf8.RuneCountInString(title)
	if titleLength < titleLengthMin || titleLength > titleLengthMax {
		return nil, errDomain.NewError("タイトルは、0文字以上50文字以内です。")
	}
	descriptionLength := utf8.RuneCountInString(description)
	if descriptionLength < titleLengthMin || descriptionLength > titleLengthMax {
		return nil, errDomain.NewError("内容は、0文字以上50文字以内です。")
	}

	return &Todo{
		todoID:           todoID,
		userID:           userID,
		title:            title,
		description:      description,
		attachedFilePath: attachedFilePath,
		completedAt:      completedAt,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		deletedAt:        deletedAt,
	}, nil
}

func Reconstruct(
	todoID int,
	userID int,
	title string,
	description string,
	attachedFilePath *string,
	completedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*Todo, error) {
	return newTodo(
		todoID,
		userID,
		title,
		description,
		attachedFilePath,
		completedAt,
		createdAt,
		updatedAt,
		deletedAt,
	)
}