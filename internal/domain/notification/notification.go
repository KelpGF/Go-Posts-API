package notification

type Notification struct {
	errors []error
}

func NewNotification() *Notification {
	return &Notification{}
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
