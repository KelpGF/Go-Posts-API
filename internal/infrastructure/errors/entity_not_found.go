package errors

type EntityNotFound struct {
	Message string
}

func NewEntityNotFound(entityName string) *EntityNotFound {
	return &EntityNotFound{
		Message: entityName + " not found",
	}
}

func (e *EntityNotFound) Error() string {
	return e.Message
}
