package repositories

import (
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"gorm.io/gorm"
)

type ListPostsRepository struct {
	db *gorm.DB
}

func NewListPostsRepository(db *gorm.DB) *ListPostsRepository {
	return &ListPostsRepository{
		db: db,
	}
}

func (r *ListPostsRepository) List(input *dto.ListPostsInput) *dto.ListPostsOutput {
	var posts []dto.Post
	query := r.db.Table("posts")

	query.Where("deleted_at IS NULL")

	if input.AuthorName != "" {
		query.Where("author_name LIKE ?", "%"+input.AuthorName+"%")
	}

	query.Order("published_at " + input.PublishedSort).Offset(input.Paginate.Offset()).Limit(input.GetLimit())

	query.Find(&posts)

	return &dto.ListPostsOutput{
		Posts: posts,
	}
}
