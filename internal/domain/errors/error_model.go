package errors

type ErrorModel struct {
	Errors  []error `json:"errors"`
	Message string  `json:"message"`
}

func NewErrorModel(errors []error, message string) *ErrorModel {
	return &ErrorModel{
		Errors:  errors,
		Message: message,
	}
}
