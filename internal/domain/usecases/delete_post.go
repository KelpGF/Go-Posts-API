package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type DeletePostUseCase interface {
	Execute(input *dto.DeletePostInput) *errors.ErrorModel
}
