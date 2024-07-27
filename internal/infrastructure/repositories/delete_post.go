package repositories

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"gorm.io/gorm"
)

type DeletePostRepository struct {
	db *gorm.DB
}

func NewDeletePostRepository(db *gorm.DB) *DeletePostRepository {
	return &DeletePostRepository{
		db: db,
	}
}

func (r *DeletePostRepository) Delete(input *dto.DeletePostInput) error {
	err := r.db.Where("id = ?", input.ID).Delete(&entities.Post{}).Error

	return err
}
