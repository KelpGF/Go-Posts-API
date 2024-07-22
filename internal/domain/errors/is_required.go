package errors

type IsRequired struct {
	Field string
}

func NewIsRequiredError(field string) IsRequired {
	return IsRequired{
		Field: field,
	}
}

func (e IsRequired) Error() string {
	return e.Field + " is required"
}
