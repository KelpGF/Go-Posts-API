package repository

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"gorm.io/gorm"
)

type CreatePostRepository struct {
	db *gorm.DB
}

func NewCreatePostRepository(db *gorm.DB) *CreatePostRepository {
	return &CreatePostRepository{
		db: db,
	}
}

func (r *CreatePostRepository) Create(input *repositories.CreatePostRepositoryInput) error {
	post := &entities.Post{
		Title:       input.Data.Title,
		Body:        input.Data.Body,
		AuthorName:  input.Data.AuthorName,
		PublishedAt: input.Data.PublishedAt,
		CreatedAt:   input.Data.CreatedAt,
	}

	if err := r.db.Create(post).Error; err != nil {
		return err
	}

	return nil
}
