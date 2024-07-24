package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type DeletePostUseCaseInput struct {
	ID string
}

type DeletePostUseCase interface {
	Execute(input *DeletePostUseCaseInput) *errors.ErrorModel
}
