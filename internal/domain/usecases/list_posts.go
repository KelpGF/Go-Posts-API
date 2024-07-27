package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
)

type ListPostsUseCase interface {
	Execute(dto.ListPostsInput) *dto.ListPostsOutput
}
