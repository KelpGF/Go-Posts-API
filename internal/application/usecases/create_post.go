package usecases

import (
	"fmt"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/models"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
)

type CreatePostUseCaseInput struct {
	Data models.CreatePost
}

type CreatePostUseCaseOutput struct {
	ID string
}

type CreatePostUseCase struct {
	CreatePostRepository repositories.CreatePostRepository
}

func NewCreatePostUseCase(createPostRepository repositories.CreatePostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{
		CreatePostRepository: createPostRepository,
	}
}

func (uc *CreatePostUseCase) Execute(input *CreatePostUseCaseInput) (*CreatePostUseCaseOutput, *errors.ErrorModel) {
	post, err := entities.NewPost(
		input.Data.Title,
		input.Data.Body,
		input.Data.AuthorName,
		input.Data.PublishedAt,
	)
	if err != nil {
		return nil, err
	}

	externalError := uc.CreatePostRepository.Create(&repositories.CreatePostRepositoryInput{
		Data: input.Data,
	})
	if externalError != nil {
		errorMessage := fmt.Sprintf("Error creating post: %s", externalError.Error())
		return nil, errors.NewErrorModel(nil, errorMessage)
	}

	return &CreatePostUseCaseOutput{ID: post.GetId()}, nil
}
