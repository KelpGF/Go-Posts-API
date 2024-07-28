package usecases

import (
	"fmt"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
)

type EditPostUseCase struct {
	findPostByIdRepository repositories.FindPostByIdRepository
	editPostRepository     repositories.EditPostRepository
}

func NewEditPostUseCase(
	findPostByIdRepository repositories.FindPostByIdRepository,
	editPostRepository repositories.EditPostRepository,
) *EditPostUseCase {
	return &EditPostUseCase{
		findPostByIdRepository: findPostByIdRepository,
		editPostRepository:     editPostRepository,
	}
}

func (uc *EditPostUseCase) Execute(input *dto.EditPostInput) *errors.ErrorModel {
	postId, err := entities.ParseID(input.ID)
	if err != nil {
		return errors.NewErrorModel(nil, "Invalid ID")
	}

	post, err := uc.findPostByIdRepository.FindById(&postId)
	if err != nil {
		errorMessage := fmt.Sprintf("Error finding post: %s", err.Error())
		return errors.NewErrorModel(nil, errorMessage)
	}

	post.SetTitle(input.Title)
	post.SetBody(input.Body)
	post.SetAuthorName(input.AuthorName)
	post.SetPublishedAt(input.PublishedAt)

	hasErrors := post.HasErrors()
	if hasErrors {
		return errors.NewErrorModel(post.GetNotificationErrors(), post.GetNotificationErrorMessage())
	}

	externalError := uc.editPostRepository.Edit(&repositories.EditPostRepositoryInput{
		Data: post,
	})
	if externalError != nil {
		errorMessage := fmt.Sprintf("Error editing post: %s", externalError.Error())
		return errors.NewErrorModel(nil, errorMessage)
	}

	return nil
}
