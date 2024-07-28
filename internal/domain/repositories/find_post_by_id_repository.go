package repositories

import (
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	entityPost "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
)

type FindPostByIdRepository interface {
	FindById(input *entities.ID) (entityPost.Post, error)
}
