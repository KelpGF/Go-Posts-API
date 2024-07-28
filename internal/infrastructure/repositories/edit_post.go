package repositories

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"gorm.io/gorm"
)

type EditPostRepository struct {
	db *gorm.DB
}

func NewEditPostRepository(db *gorm.DB) *EditPostRepository {
	return &EditPostRepository{
		db: db,
	}
}

func (r *EditPostRepository) Edit(input *repositories.EditPostRepositoryInput) error {
	post := entities.Post{}
	post.ID = input.Data.GetId()
	post.Title = input.Data.GetTitle()
	post.Body = input.Data.GetBody()
	post.AuthorName = input.Data.GetAuthorName()
	post.PublishedAt = input.Data.GetPublishedAt()

	return r.db.Model(&post).Updates(post).Error
}
