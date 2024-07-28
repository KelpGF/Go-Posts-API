package repositories

import (
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
)

type EditPostRepositoryInput struct {
	Data entities.Post
}

type EditPostRepository interface {
	Edit(input *EditPostRepositoryInput) error
}
