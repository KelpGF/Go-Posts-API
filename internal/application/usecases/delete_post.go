package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
)

type DeletePostUseCase struct {
	deletePostRepository repositories.DeletePostRepository
}

func NewDeletePostUseCase(deletePostRepository repositories.DeletePostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{
		deletePostRepository: deletePostRepository,
	}
}

func (uc *DeletePostUseCase) Execute(input *dto.DeletePostInput) *errors.ErrorModel {
	deletePostRepositoryInput := &repositories.DeletePostRepositoryInput{
		ID: input.ID,
	}

	err := uc.deletePostRepository.Delete(deletePostRepositoryInput)
	if err != nil {
		return errors.NewErrorModel(nil, "We couldn't delete the post! Check if the ID is correct.")
	}

	return nil
}
