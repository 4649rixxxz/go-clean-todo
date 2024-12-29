package usecase

/*
Usecase層のエラーを定義する。
目的はプレゼンテーション層などに依存しないエラーを定義することで、「usecase層のエラーをどのように扱うか」はエラーを使用する側に委ねるため。
*/

const (
	InvalidInputError     = "INVALID_INPUT"
	ResourceNotFoundError = "RESOURCE_NOT_FOUND"
	InternalServerError   = "INTERNAL_SERVER_ERROR"
)

type UsecaseErrorI interface {
	Code() string
	Field() string
	Error() string
}

type UsecaseError struct {
	code    string
	field   *string
	message string
}

func (e *UsecaseError) Code() string {
	return e.code
}

func (e *UsecaseError) Field() string {
	return *e.field
}

func (e *UsecaseError) Error() string {
	return e.message
}

func NewInvalidInputError(field, message string) *UsecaseError {
	return &UsecaseError{code: InvalidInputError, field: &field, message: message}
}
func NewResourceNotFoundError(field, message string) *UsecaseError {
	return &UsecaseError{code: ResourceNotFoundError, field: &field, message: message}
}
func NewInternalServerError(message string) *UsecaseError {
	return &UsecaseError{code: InternalServerError, message: message}
}
