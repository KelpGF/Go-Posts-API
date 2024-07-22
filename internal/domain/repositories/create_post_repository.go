package repositories

import "github.com/KelpGF/Go-Posts-API/internal/domain/models"

type CreatePostRepositoryInput struct {
	Data models.CreatePost
}

type CreatePostRepository interface {
	Create(input *CreatePostRepositoryInput) error
}
