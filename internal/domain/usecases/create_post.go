package usecases

import (
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type CreatePostUseCaseInput struct {
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	AuthorName  string    `json:"author_name"`
	PublishedAt time.Time `json:"published_at"`
}

type CreatePostUseCaseOutput struct {
	ID string
}

type CreatePostUseCase interface {
	Execute(input *CreatePostUseCaseInput) (*CreatePostUseCaseOutput, *errors.ErrorModel)
}
