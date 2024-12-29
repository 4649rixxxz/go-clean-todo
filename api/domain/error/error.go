package error

type DomainErrorI interface {
	Field() string
	Error() string
}

type DomainError struct {
	field   string
	message string
}

func (e *DomainError) Field() string {
	return e.field
}

func (e *DomainError) Error() string {
	return e.message
}

func NewDomainError(field, message string) DomainErrorI {
	return &DomainError{
		field:   field,
		message: message,
	}
}
