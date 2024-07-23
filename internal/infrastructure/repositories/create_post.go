package repositories

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
		ID:          input.Data.GetId(),
		Title:       input.Data.GetTitle(),
		Body:        input.Data.GetBody(),
		AuthorName:  input.Data.GetAuthorName(),
		PublishedAt: input.Data.GetPublishedAt(),
		CreatedAt:   input.Data.GetCreatedAt(),
	}

	if err := r.db.Create(post).Error; err != nil {
		return err
	}

	return nil
}
