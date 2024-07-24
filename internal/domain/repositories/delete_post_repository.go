package repositories

type DeletePostRepositoryInput struct {
	ID string
}

type DeletePostRepository interface {
	Delete(input *DeletePostRepositoryInput) error
}
