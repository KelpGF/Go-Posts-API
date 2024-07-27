package usecases

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
)

type ListPostsUseCase struct {
	listPostsRepository repositories.ListPostsRepository
}

func NewListPostsUseCase(listPostsRepository repositories.ListPostsRepository) *ListPostsUseCase {
	return &ListPostsUseCase{
		listPostsRepository: listPostsRepository,
	}
}

func (uc *ListPostsUseCase) Execute(input *dto.ListPostsInput) *dto.ListPostsOutput {
	inputRepository := &dto.ListPostsInput{
		AuthorName:    input.AuthorName,
		PublishedSort: input.PublishedSort,
		Paginate:      input.Paginate,
	}

	if input.PublishedSort != "asc" && input.PublishedSort != "desc" {
		inputRepository.PublishedSort = "desc"
	}

	output := uc.listPostsRepository.List(inputRepository)

	return output
}
