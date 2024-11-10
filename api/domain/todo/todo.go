package todo

import (
	"time"
	"unicode/utf8"

	errDomain "go-clean-todo/domain/error"
)

type Todo struct {
	todoID           uint
	userID           uint
	title            string
	description      string
	attachedFilePath *string
	completedAt      *time.Time
	createdAt        time.Time
	updatedAt        time.Time
}

const (
	titleLengthMin       = 0
	titleLengthMax       = 50
	descriptionLengthMin = 0
	descriptionLengthMax = 200
)

func newTodo(
	todoID uint,
	userID uint,
	title string,
	description string,
	attachedFilePath *string,
	completedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
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
	}, nil
}

func NewTodo(
	userID uint,
	title string,
	description string,
) (*Todo, error) {
	return newTodo(
		0,
		userID,
		title,
		description,
		nil,
		nil,
		time.Now(),
		time.Now(),
	)
}

func Reconstruct(
	todoID uint,
	userID uint,
	title string,
	description string,
	attachedFilePath *string,
	completedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
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
	)
}

func (t *Todo) TodoID() uint {
	return t.todoID
}

func (t *Todo) UserID() uint {
	return t.userID
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Description() string {
	return t.description
}
func (t *Todo) AttachedFilePath() *string {
	return t.attachedFilePath
}

func (t *Todo) CompletedAt() *time.Time {
	return t.completedAt
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) UpdatedAt() time.Time {
	return t.updatedAt
}
