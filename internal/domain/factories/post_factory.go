package factories

import (
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
)

type PostFactory interface {
	NewPost(title string, body string, authorName string, publishedAt time.Time) (entities.Post, *errors.ErrorModel)
}

type RestorePostFactory interface {
	Restore(id string, title string, body string, authorName string, publishedAt time.Time, createdAt time.Time) entities.Post
}
