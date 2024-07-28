package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type EditPostUseCase interface {
	Execute(input *dto.EditPostInput) *errors.ErrorModel
}
