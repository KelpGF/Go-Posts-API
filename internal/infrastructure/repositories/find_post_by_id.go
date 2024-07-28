package repositories

import (
	"gorm.io/gorm"

	entityId "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	entityPost "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/domain/factories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/errors"
)

type FindPostByIdRepository struct {
	db *gorm.DB

	postFactory factories.RestorePostFactory
}

func NewFindPostByIdRepository(
	db *gorm.DB,
	postFactory factories.RestorePostFactory,
) *FindPostByIdRepository {
	return &FindPostByIdRepository{
		db:          db,
		postFactory: postFactory,
	}
}

func (r *FindPostByIdRepository) FindById(input *entityId.ID) (entityPost.Post, error) {
	var post entities.Post
	r.db.First(&post, "id = ?", input)

	if post.ID == "" {
		return nil, errors.NewEntityNotFound("Post")
	}

	restoredPost := r.postFactory.Restore(
		post.ID,
		post.Title,
		post.Body,
		post.AuthorName,
		post.PublishedAt.Local(),
	)

	return restoredPost, nil
}
