package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type CreatePostUseCase interface {
	Execute(input *dto.CreatePostInput) (*dto.CreatePostOutput, *errors.ErrorModel)
}
