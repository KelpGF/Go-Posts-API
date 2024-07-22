package notification

type Notification struct {
	context string
	errors  []error
}

func NewNotification(context string) *Notification {
	return &Notification{
		context: context,
	}
}

func (n *Notification) AddError(err error) {
	n.errors = append(n.errors, err)
}

func (n *Notification) GetErrors() []error {
	return n.errors
}

func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

func (n *Notification) GetErrorsMessage() string {
	var errorsString string = n.context + ": "

	length := len(n.errors)

	for idx, err := range n.errors {
		errorsString += err.Error()

		if idx < length-1 {
			errorsString += ", "
		}
	}

	return errorsString
}
