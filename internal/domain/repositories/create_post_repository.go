package repositories

import (
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
)

type CreatePostRepositoryInput struct {
	Data entities.Post
}

type CreatePostRepository interface {
	Create(input *CreatePostRepositoryInput) error
}
