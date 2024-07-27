package usecases

import (
	"fmt"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/factories"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
)

type CreatePostUseCase struct {
	CreatePostRepository repositories.CreatePostRepository
	PostFactory          factories.PostFactory
}

func NewCreatePostUseCase(
	createPostRepository repositories.CreatePostRepository,
	postFactory factories.PostFactory,
) *CreatePostUseCase {
	return &CreatePostUseCase{
		CreatePostRepository: createPostRepository,
		PostFactory:          postFactory,
	}
}

func (uc *CreatePostUseCase) Execute(input *dto.CreatePostInput) (*dto.CreatePostOutput, *errors.ErrorModel) {
	post, err := uc.PostFactory.NewPost(
		input.Title,
		input.Body,
		input.AuthorName,
		input.PublishedAt,
	)

	if err != nil {
		return nil, err
	}

	externalError := uc.CreatePostRepository.Create(&repositories.CreatePostRepositoryInput{
		Data: post,
	})
	if externalError != nil {
		errorMessage := fmt.Sprintf("Error creating post: %s", externalError.Error())
		return nil, errors.NewErrorModel(nil, errorMessage)
	}

	return &dto.CreatePostOutput{ID: post.GetId()}, nil
}
