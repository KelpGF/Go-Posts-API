package errors

type IsRequired struct {
	Field   string `json:"-"`
	Message string `json:"message"`
}

func NewIsRequiredError(field string) IsRequired {
	return IsRequired{
		Field:   field,
		Message: field + " is required",
	}
}

func (e IsRequired) Error() string {
	return e.Message
}
